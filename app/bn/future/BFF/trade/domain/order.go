package domain

import bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"

type Order struct {
	AccumID      string
	MaxAccum     string
	Accum        string
	ClientId     string
	Symbol       string
	PositionSide string
	Side         string
	AmountB      string
	AccountId    string
	CreatedAt    string
	OrderStatus  string
}

func (o *Order) NewOrderFrom(d *Order) *Order {
	return &Order{
		AccumID:      d.AccumID,
		MaxAccum:     d.MaxAccum,
		Accum:        d.Accum,
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.AmountB,
		AccountId:    d.AccountId,
		CreatedAt:    d.CreatedAt,
		OrderStatus:  d.OrderStatus,
	}
}

func (o *Order) IsBuyPosition() bool {
	if o.Side == bnconstant.BUY && o.PositionSide == bnconstant.LONG {
		return true
	}
	if o.Side == bnconstant.SELL && o.PositionSide == bnconstant.SHORT {
		return true
	}
	return false
}
