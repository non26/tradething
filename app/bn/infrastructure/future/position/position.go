package infrastructure

import (
	"context"
	req "tradething/app/bn/infrastructure/adaptor/future/req/future_trade"

	positionconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/non26/tradepkg/pkg/bn/utils"
	"github.com/shopspring/decimal"
)

type IPosition interface {
	BuyPosition(ctx context.Context, position *Position) error
	SellPosition(ctx context.Context, position *Position) error
}

type Position struct {
	PositionSide string
	AmountB      string
	Symbol       string
	OrderType    string
	ClientId     string
	Side         string
}

func (p *Position) ToPlacePositionModel() *req.PlacePosition {
	return &req.PlacePosition{
		PositionSide:  p.PositionSide,
		Side:          p.Side,
		EntryQuantity: p.AmountB,
		Symbol:        p.Symbol,
		ClientOrderId: p.ClientId,
	}
}

func (p *Position) IsLongPosition() bool {
	return p.PositionSide == positionconstant.LONG
}

func (p *Position) IsShortPosition() bool {
	return p.PositionSide == positionconstant.SHORT
}

func (p *Position) SetDefaultClientId(counting int) {
	if p.ClientId == "" {
		p.ClientId = utils.BinanceDefaultClientID(p.Symbol, p.PositionSide, counting)
	}
}

func (p *Position) AddMoreAmountB(amountB string) error {
	amountQInt, err := decimal.NewFromString(amountB)
	if err != nil {
		return err
	}
	prevAmountQInt, err := decimal.NewFromString(p.AmountB)
	if err != nil {
		return err
	}
	p.AmountB = amountQInt.Add(prevAmountQInt).String()
	return nil
}
