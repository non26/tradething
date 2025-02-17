package req

import "tradething/app/bn/internal/process/spot/domain"

type Trade struct {
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	AmountB  string `json:"amountB"`
	ClientId string `json:"clientId"`
}

func (t *Trade) ToOrder() *domain.Order {
	return &domain.Order{
		Symbol:   t.Symbol,
		Side:     t.Side,
		AmountB:  t.AmountB,
		ClientId: t.ClientId,
	}
}
