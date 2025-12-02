package res

import "tradething/app/bn/future/BFF/trade/domain"

type AccumulateOrderMultiRes struct {
	Data []AccumulateOrderRes `json:"data"`
}

func (r *AccumulateOrderMultiRes) FromDomain(orders []*domain.Order) {
	r.Data = make([]AccumulateOrderRes, 0)
	for _, order := range orders {
		orderRes := AccumulateOrderRes{}
		orderRes.FromDomain(order)
		r.Data = append(r.Data, orderRes)
	}
}
