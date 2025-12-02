package req

import "tradething/app/bn/future/BFF/trade/domain"

type AccumulateOrderMultiReq struct {
	Data []AccumulateOrderReq `json:"data"`
}

func (r *AccumulateOrderMultiReq) ToDomain() []*domain.Order {
	orders := make([]*domain.Order, 0)
	for _, item := range r.Data {
		order := item.ToDomain()
		orders = append(orders, order)
	}
	return orders
}
