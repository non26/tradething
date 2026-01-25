package dto

import (
	"tradething/app/bn/future/BFF/position_mangement/domain"
	externalpsotiionhistoryDomain "tradething/app/bn/future/position_history/domain"
)

type BnFtHistory struct {
	ClientId     string
	Symbol       string
	PositionSide string
	AccountId    string
}

func NewBnFtHistoryDto() *BnFtHistory {
	return &BnFtHistory{}
}

func (b *BnFtHistory) FromDomainToExternalPositionHistoryDomain(d *domain.Position) *externalpsotiionhistoryDomain.BnFtHistory {
	return &externalpsotiionhistoryDomain.BnFtHistory{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		AccountId:    d.AccountId,
	}
}

func (b *BnFtHistory) FromExternalPositionHistoryDomainToDomain(d *externalpsotiionhistoryDomain.BnFtHistory) *domain.Position {
	return &domain.Position{
		ClientId:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		AccountId:    d.AccountId,
	}
}
