package domain

import positionExt "tradething/app/bn/future/position/domain"

type Position struct {
	Symbol       string
	PositionSide string
	Side         string
	AmountB      string
	AccountId    string
	ClientId     string
}

func (p *Position) ToOpeningPositionExtSrv() *positionExt.Position {
	return &positionExt.Position{
		ClientId:     p.ClientId,
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
		Side:         p.Side,
		AmountB:      p.AmountB,
		AccountId:    p.AccountId,
	}
}
