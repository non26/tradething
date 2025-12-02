package res

import "tradething/app/bn/future/BFF/trade/domain"

type ManageOrderRes struct {
	CloseOrderById  CloseOrderByIdMultiRes  `json:"close_order_by_id"`
	NewOrder        NewOrderMultiRes        `json:"new_order"`
	AccumulateOrder AccumulateOrderMultiRes `json:"accumulate_order"`
}

func (r *ManageOrderRes) FromDomain(orders []*domain.Order) {
	r.CloseOrderById.FromDomain(orders)
	r.NewOrder.FromDomain(orders)
	r.AccumulateOrder.FromDomain(orders)
}
