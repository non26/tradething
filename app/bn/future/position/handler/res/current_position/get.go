package res

import "tradething/app/bn/future/position/domain"

type GetCurrentPositionRes struct {
	ClientId     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	Side         string `json:"side"`
	AmountB      string `json:"amount_b"`
	AccountId    string `json:"account_id"`
}

func (r *GetCurrentPositionRes) FromDomain(position *domain.Position) *GetCurrentPositionRes {
	return &GetCurrentPositionRes{
		ClientId:     position.ClientId,
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		Side:         position.Side,
		AmountB:      position.AmountB,
		AccountId:    position.AccountId,
	}
}
