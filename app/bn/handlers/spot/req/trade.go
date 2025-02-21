package req

import "tradething/app/bn/process/spot/domain"

type Trade struct {
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	AmountB  string `json:"amount_b"`
	ClientId string `json:"client_id"`
}

func (t *Trade) ToOrder() *domain.Order {
	return &domain.Order{
		Symbol:   t.Symbol,
		Side:     t.Side,
		AmountB:  t.AmountB,
		ClientId: t.ClientId,
	}
}
