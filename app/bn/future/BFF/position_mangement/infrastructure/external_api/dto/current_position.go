package dto

import (
	"tradething/app/bn/future/BFF/position_mangement/domain"
	posiitionDoamin "tradething/app/bn/future/position/domain"
)

type CurrentPositionDto struct {
	Symbol       string
	PositionSide string
	Side         string
	AmountB      string
	AccountId    string
	ClientId     string
}

func NewCurrentPositionDto() *CurrentPositionDto {
	return &CurrentPositionDto{}
}

func (c *CurrentPositionDto) FromDomain(domain *domain.Position) *posiitionDoamin.Position {
	return &posiitionDoamin.Position{
		Symbol:       domain.Symbol,
		PositionSide: domain.PositionSide,
		Side:         domain.Side,
		AmountB:      domain.AmountB,
		AccountId:    domain.AccountId,
		ClientId:     domain.ClientId,
	}
}

func (c *CurrentPositionDto) ToDomain(d *posiitionDoamin.Position) *domain.Position {
	return &domain.Position{
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.AmountB,
		AccountId:    d.AccountId,
		ClientId:     d.ClientId,
	}
}
