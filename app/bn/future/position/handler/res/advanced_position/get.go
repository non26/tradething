package res

import "tradething/app/bn/future/position/domain"

type GetAdvancedPositionRes struct {
	ClientId     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	Side         string `json:"side"`
	AmountB      string `json:"amount_b"`
	AccountId    string `json:"account_id"`
}

func (r *GetAdvancedPositionRes) FromDomain(position *domain.Position) *GetAdvancedPositionRes {
	return &GetAdvancedPositionRes{
		ClientId:     position.ClientId,
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		Side:         position.Side,
		AmountB:      position.AmountB,
		AccountId:    position.AccountId,
	}
}
