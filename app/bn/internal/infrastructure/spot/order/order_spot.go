package order

import (
	"context"
	"errors"
	spot "tradething/app/bn/internal/infrastructure/adaptor/spot"
	spotreq "tradething/app/bn/internal/infrastructure/adaptor/spot/req"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_spot/models"
	"github.com/non26/tradepkg/pkg/bn/utils"
	"github.com/shopspring/decimal"
)

type Order struct {
	Symbol           string
	Side             string
	Type             string
	Quantity         string
	NewClientOrderId string
}

func (o *Order) SetDefaultClientId(counting int) {
	if o.NewClientOrderId == "" {
		o.NewClientOrderId = utils.BinanceDefaultClientID(o.Symbol, "post", counting)
	}
}

func (o *Order) ToOpeningSpotTable() *dynamodbmodel.BnSpotOpeningPosition {
	return &dynamodbmodel.BnSpotOpeningPosition{
		ClientId: o.NewClientOrderId,
		Symbol:   o.Symbol,
		AmountB:  o.Quantity,
	}
}

func (o *Order) ToSpotOrderRequest() *spotreq.SpotOrderRequest {
	return &spotreq.SpotOrderRequest{
		Symbol:           o.Symbol,
		Side:             o.Side,
		Type:             o.Type,
		Quantity:         o.Quantity,
		NewClientOrderId: o.NewClientOrderId,
	}
}

func (o *Order) AddMoreAmountB(amountB string) error {
	amountQInt, err := decimal.NewFromString(amountB)
	if err != nil {
		return err
	}
	prevAmountQInt, err := decimal.NewFromString(o.Quantity)
	if err != nil {
		return err
	}
	o.Quantity = amountQInt.Add(prevAmountQInt).String()
	return nil
}

func (o *Order) ToHistoryTable() *dynamodbmodel.BnSpotHistory {
	return &dynamodbmodel.BnSpotHistory{
		ClientId: o.NewClientOrderId,
		Symbol:   o.Symbol,
	}
}

type OrderSpot struct {
	adaptor                    spot.IBinanceSpotTradeService
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewOrderSpot(
	adaptor spot.IBinanceSpotTradeService,
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) *OrderSpot {
	return &OrderSpot{
		adaptor:                    adaptor,
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (o *OrderSpot) BuyOrder(ctx context.Context, order Order) error {
	crypto, err := o.bnSpotCryptoTable.Get(context.Background(), order.Symbol)
	if err != nil {
		return err
	}

	if !crypto.IsFound() {
		crypto = dynamodbmodel.NewBinanceSpotCryptoTableRecord(order.Symbol)
	}
	order.SetDefaultClientId(crypto.GetCounting())

	openingSpot, err := o.bnSpotOpeningPositionTable.Get(context.Background(), order.ToOpeningSpotTable())
	if err != nil {
		return err
	}

	if openingSpot.ClientId == order.NewClientOrderId {
		return errors.New("duplicate client id")
	}

	_, err = o.adaptor.PlaceOrder(context.Background(), order.ToSpotOrderRequest())
	if err != nil {
		return err
	}

	if openingSpot.IsFound() {
		err = order.AddMoreAmountB(order.Quantity)
		if err != nil {
			return err
		}
	}

	err = o.bnSpotOpeningPositionTable.Upsert(context.Background(), order.ToOpeningSpotTable())
	if err != nil {
		return err
	}

	err = o.bnSpotCryptoTable.Update(context.Background(), crypto)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderSpot) SellOrder(ctx context.Context, order Order) error {

	_, err := o.adaptor.PlaceOrder(context.Background(), order.ToSpotOrderRequest())
	if err != nil {
		return err
	}

	err = o.bnSpotOpeningPositionTable.Delete(context.Background(), order.ToOpeningSpotTable())
	if err != nil {
		return err
	}

	err = o.bnSpotHistoryTable.Insert(context.Background(), order.ToHistoryTable())
	if err != nil {
		return err
	}

	return nil
}
