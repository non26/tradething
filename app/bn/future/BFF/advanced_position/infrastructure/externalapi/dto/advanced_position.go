package dto

import (
	"tradething/app/bn/future/BFF/advanced_position/domain"
	posiitionDoamin "tradething/app/bn/future/position/domain"
)

type AdvancedPositionDto struct {
	ClientId     string
	Symbol       string
	PositionSide string
	Side         string
	AmountB      string
	AccountId    string
}

func NewAdvancedPositionDto() *AdvancedPositionDto {
	return &AdvancedPositionDto{}
}

func (a *AdvancedPositionDto) FromDomain(d *domain.AdvancedPosition) *posiitionDoamin.Position {
	return &posiitionDoamin.Position{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.AmountB,
		AccountId:    d.AccountId,
	}
}

func (a *AdvancedPositionDto) ToDomain(d *posiitionDoamin.Position) *domain.AdvancedPosition {
	return &domain.AdvancedPosition{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.AmountB,
		AccountId:    d.AccountId,
	}
}
