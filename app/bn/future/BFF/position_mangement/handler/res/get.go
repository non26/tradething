package res

import "tradething/app/bn/future/BFF/position_mangement/domain"

type GetRes struct {
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
	PositionSide string `json:"positionSide"`
	Side         string `json:"side"`
	AmountB      string `json:"amountB"`
	ClientId     string `json:"clientId"`
}

func (r *GetRes) FromDomain(position *domain.Position) *GetRes {
	return &GetRes{
		Symbol:       position.Symbol,
		AccountId:    position.AccountId,
		PositionSide: position.PositionSide,
		Side:         position.Side,
		AmountB:      position.AmountB,
		ClientId:     position.ClientId,
	}
}
