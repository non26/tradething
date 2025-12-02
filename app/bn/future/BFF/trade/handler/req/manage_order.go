package req

import "tradething/app/bn/future/BFF/trade/domain"

type ManageOrderReq struct {
	CloseClientID   CloseOrderByIdMultiReq  `json:"close_client_id"`
	NewOrder        NewOrderMultiReq        `json:"new_order"`
	AccumulateOrder AccumulateOrderMultiReq `json:"accumulate_order"`
}

func (r *ManageOrderReq) ToCloseOrderByIdDomain() []*domain.Order {
	return r.CloseClientID.ToDomain()
}

func (r *ManageOrderReq) ToNewOrderDomain() []*domain.Order {
	return r.NewOrder.ToDomain()
}

func (r *ManageOrderReq) ToAccumulateOrderDomain() []*domain.Order {
	return r.AccumulateOrder.ToDomain()
}
