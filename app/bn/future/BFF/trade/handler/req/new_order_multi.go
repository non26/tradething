package req

import "tradething/app/bn/future/BFF/trade/domain"

type NewOrderMultiReq struct {
	Data []NewOrderReq `json:"data"`
}

func (r *NewOrderMultiReq) ToDomain() []*domain.Order {
	orders := make([]*domain.Order, 0)
	for _, item := range r.Data {
		order := item.ToDomain()
		orders = append(orders, order)
	}
	return orders
}
