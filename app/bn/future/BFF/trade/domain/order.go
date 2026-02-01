package domain

import (
	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/shopspring/decimal"
)

type Order struct {
	AccumID      string
	MaxAccum     string // max accum amount
	Accum        string // accum amount
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

func (o *Order) ToSellPosition() {
	if o.PositionSide == bnconstant.LONG {
		o.Side = bnconstant.SELL
	} else {
		o.Side = bnconstant.BUY
	}
}

func (o *Order) IsAmount1ExceedAmount2(amount1 string, amount2 string) (bool, error) {
	_amount1, err := decimal.NewFromString(amount1)
	if err != nil {
		return false, err
	}
	_amount2, err := decimal.NewFromString(amount2)
	if err != nil {
		return false, err
	}
	return _amount1.GreaterThanOrEqual(_amount2), nil
}

func (o *Order) AddAmount(amount string, addBy string) (string, error) {
	_amount, err := decimal.NewFromString(amount)
	if err != nil {
		return "", err
	}
	_addBy, err := decimal.NewFromString(addBy)
	if err != nil {
		return "", err
	}
	return _amount.Add(_addBy).String(), nil
}

func (o *Order) IsAdvancedPosition() bool {
	if o.ClientId != "" && (o.Symbol == "" &&
		o.AccountId == "" &&
		o.PositionSide == "" &&
		o.Side == "" &&
		o.AmountB == "") {
		return true
	}
	return false
}
