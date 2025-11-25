package domain

import "tradething/app/bn/future/position/domain"

type AdvancedPosition struct {
	ClientId     string
	Symbol       string
	PositionSide string
	Side         string
	AmountB      string
	AccountId    string
}

func (p *AdvancedPosition) ToPositionExtSrv() *domain.Position {
	return &domain.Position{
		ClientId:     p.ClientId,
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
		Side:         p.Side,
		AmountB:      p.AmountB,
		AccountId:    p.AccountId,
	}
}
