package res

import "tradething/app/bn/future/BFF/trade/domain"

type NewOrderMultiRes struct {
	Data []NewOrderRes `json:"data"`
}

func (r *NewOrderMultiRes) FromDomain(orders []*domain.Order) {
	r.Data = make([]NewOrderRes, 0)
	for _, order := range orders {
		orderRes := NewOrderRes{}
		orderRes.FromDomain(order)
		r.Data = append(r.Data, orderRes)
	}
}
