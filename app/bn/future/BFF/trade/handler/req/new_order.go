package req

import "tradething/app/bn/future/BFF/trade/domain"

type NewOrderReq struct {
	ClientID     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	Side         string `json:"side"`
	AmountB      string `json:"amount_b"`
	AccountId    string `json:"account_id"`
}

func (r *NewOrderReq) ToDomain() *domain.Order {
	return &domain.Order{
		ClientId:     r.ClientID,
		Symbol:       r.Symbol,
		PositionSide: r.PositionSide,
		Side:         r.Side,
		AmountB:      r.AmountB,
		AccountId:    r.AccountId,
	}
}
