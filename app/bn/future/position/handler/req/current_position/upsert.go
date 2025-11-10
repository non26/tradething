package req

import "tradething/app/bn/future/position/domain"

type UpsertCurrentPositionReq struct {
	ClientId     string `json:"client_id" binding:"required"`
	Symbol       string `json:"symbol" binding:"required"`
	PositionSide string `json:"position_side" binding:"required"`
	Side         string `json:"side" binding:"required"`
	AmountB      string `json:"amount_b" binding:"required"`
	AccountId    string `json:"account_id" binding:"required"`
}

func (r *UpsertCurrentPositionReq) ToDomain() *domain.Position {
	return &domain.Position{
		ClientId:     r.ClientId,
		Symbol:       r.Symbol,
		PositionSide: r.PositionSide,
		Side:         r.Side,
		AmountB:      r.AmountB,
		AccountId:    r.AccountId,
	}
}
