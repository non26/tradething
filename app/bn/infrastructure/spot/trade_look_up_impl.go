package infrastructure

import (
	"context"
	"errors"
	"tradething/app/bn/infrastructure/spot/order"
	domainservice "tradething/app/bn/process/spot/domain_service/trade"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_spot/models"
)

type tradeLookUp struct {
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewTradeLookUp(
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ITradeLookUp {
	return &tradeLookUp{
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (t *tradeLookUp) LookUp(ctx context.Context, order *order.Order) (*domainservice.TradeLookUp, error) {

	spotHistory, err := t.bnSpotHistoryTable.Get(ctx, order.NewClientOrderId)
	if err != nil {
		return nil, err
	}

	if spotHistory.IsFound() {
		return nil, errors.New("duplicate client id")
	}

	crypto, err := t.bnSpotCryptoTable.Get(ctx, order.Symbol)
	if err != nil {
		return nil, err
	}

	if !crypto.IsFound() {
		crypto = dynamodbmodel.NewBinanceSpotCryptoTableRecord(order.Symbol)
	} else {
		crypto.SetCounting(crypto.GetCounting() + 1)
	}
	order.SetDefaultClientId(crypto.GetCounting())

	openingSpot, err := t.bnSpotOpeningPositionTable.Get(ctx, order.ToOpeningSpotTable())
	if err != nil {
		return nil, err
	}

	if order.Side == bnconstant.BUY {
		if openingSpot.ClientId == order.NewClientOrderId {
			return nil, errors.New("duplicate client id")
		}
	}

	lookup := domainservice.NewTradeLookUp()
	lookup.OpeningPosition.SetIsFound(openingSpot.IsFound())
	lookup.OpeningPosition.SetClientId(openingSpot.ClientId)
	lookup.OpeningPosition.SetSymbol(order.Symbol)
	lookup.OpeningPosition.SetQuantity(order.Quantity)
	lookup.Symbol.SetSymbol(order.Symbol)
	lookup.Symbol.SetCounting(crypto.GetCounting())
	return lookup, nil
}
