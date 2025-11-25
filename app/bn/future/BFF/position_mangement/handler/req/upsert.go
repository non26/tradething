package req

import "tradething/app/bn/future/BFF/position_mangement/domain"

type UpsertReq struct {
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
	PositionSide string `json:"positionSide"`
	Side         string `json:"side"`
	AmountB      string `json:"amountB"`
	ClientId     string `json:"clientId"`
}

func (r *UpsertReq) ToDomain() *domain.Position {
	return &domain.Position{
		Symbol:       r.Symbol,
		AccountId:    r.AccountId,
		PositionSide: r.PositionSide,
		Side:         r.Side,
		AmountB:      r.AmountB,
		ClientId:     r.ClientId,
	}
}
