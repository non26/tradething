package res

import "tradething/app/bn/future/BFF/trade/domain"

type CloseOrderByIdMultiRes struct {
	Data []CloseOrderByIdRes `json:"data"`
}

func (r *CloseOrderByIdMultiRes) FromDomain(orders []*domain.Order) {
	r.Data = make([]CloseOrderByIdRes, 0)
	for _, order := range orders {
		orderRes := CloseOrderByIdRes{}
		orderRes.FromDomain(order)
		r.Data = append(r.Data, orderRes)
	}
}
