package bnfuture

type CloseByClientIdServiceRequest struct {
	OrderIds []string `json:"order_ids"`
}

func (m *CloseByClientIdServiceRequest) GetOrderIds() []string {
	return m.OrderIds
}

func (m *CloseByClientIdServiceRequest) SetOrderIds(orderIds []string) {
	m.OrderIds = orderIds
}
