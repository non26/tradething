package req

import "tradething/app/bn/process/spot/domain"

type MultipleOrders struct {
	Orders []Trade `json:"orders"`
}

func (m *MultipleOrders) ToDomain() []domain.Order {
	orders := make([]domain.Order, len(m.Orders))
	for i, order := range m.Orders {
		orders[i] = order.ToDomain()
	}
	return orders
}

func (t *Trade) ToDomain() domain.Order {
	return domain.Order{
		Symbol:   t.Symbol,
		Side:     t.Side,
		AmountB:  t.AmountB,
		ClientId: t.ClientId,
	}
}
