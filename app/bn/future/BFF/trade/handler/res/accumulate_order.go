package res

import "tradething/app/bn/future/BFF/trade/domain"

type AccumulateOrderRes struct {
	AccumID           string `json:"accum_id"`
	ReferenceClientID string `json:"reference_client_id"`
	MaxAccum          string `json:"max_accum"`
	Accum             string `json:"accum"`
	OrderStatus       string `json:"order_status,omitempty"`
}

func (r *AccumulateOrderRes) FromDomain(order *domain.Order) {
	r.AccumID = order.AccumID
	r.ReferenceClientID = order.ClientId
	r.MaxAccum = order.MaxAccum
	r.Accum = order.Accum
	r.OrderStatus = order.OrderStatus
}
