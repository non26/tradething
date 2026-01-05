package dto

import (
	"tradething/app/bn/future/BFF/trade/domain"
	positonHistoryDomain "tradething/app/bn/future/position_history/domain"
)

type PositionHistory struct {
	ClientId     string
	Symbol       string
	PositionSide string
	AccountId    string
}

func NewPositionHistoryDto() *PositionHistory {
	return &PositionHistory{}
}

func (p *PositionHistory) ToDomain(d *positonHistoryDomain.BnFtHistory) *domain.Order {
	if d == nil {
		return nil
	}

	return &domain.Order{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		AccountId:    d.AccountId,
	}
}

func (p *PositionHistory) FromDomain(d *domain.Order) *positonHistoryDomain.BnFtHistory {
	return &positonHistoryDomain.BnFtHistory{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		AccountId:    d.AccountId,
	}
}
