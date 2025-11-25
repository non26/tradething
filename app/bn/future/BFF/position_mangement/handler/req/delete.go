package req

import "tradething/app/bn/future/BFF/position_mangement/domain"

type DeleteReq struct {
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
	PositionSide string `json:"positionSide"`
}

func (r *DeleteReq) ToDomain() *domain.Position {
	return &domain.Position{
		Symbol:       r.Symbol,
		AccountId:    r.AccountId,
		PositionSide: r.PositionSide,
	}
}
