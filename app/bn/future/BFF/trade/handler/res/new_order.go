package res

import "tradething/app/bn/future/BFF/trade/domain"

type NewOrderRes struct {
	ClientId    string `json:"client_id"`
	OrderStatus string `json:"order_status,omitempty"`
}

func (r *NewOrderRes) FromDomain(order *domain.Order) {
	r.ClientId = order.ClientId
	r.OrderStatus = order.OrderStatus
}
