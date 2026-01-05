package dto

import (
	"tradething/app/bn/future/BFF/trade/domain"
	positonDomain "tradething/app/bn/future/position/domain"
)

type Position struct {
	ClientId              string
	Symbol                string
	AccountId             string
	PositionSide          string
	Side                  string
	AmountB               string
	CreatedAt             string
	MaxAccumulatePosition string
}

func NewPositionDto() *Position {
	return &Position{}
}

func (p *Position) ToDomain(d *positonDomain.Position) *domain.Order {
	if d == nil {
		return nil
	}

	return &domain.Order{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		AccountId:    d.AccountId,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.AmountB,
	}
}

func (p *Position) FromDomain(d *domain.Order) *positonDomain.Position {
	return &positonDomain.Position{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		AccountId:    d.AccountId,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.AmountB,
	}
}
