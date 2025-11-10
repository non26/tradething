package req

import "tradething/app/bn/future/position_history/domain"

type InsertReq struct {
	ClientId     string `json:"client_id" binding:"required"`
	Symbol       string `json:"symbol" binding:"required"`
	PositionSide string `json:"position_side" binding:"required"`
	AccountId    string `json:"account_id" binding:"required"`
}

func (r *InsertReq) ToDomain() *domain.BnFtHistory {
	return &domain.BnFtHistory{
		ClientId:     r.ClientId,
		Symbol:       r.Symbol,
		PositionSide: r.PositionSide,
		AccountId:    r.AccountId,
	}
}
