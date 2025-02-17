package res

type MultipleOrder struct {
	Orders []MultipleOrderResponse
}

func (m *MultipleOrder) Add(order MultipleOrderResponse) {
	m.Orders = append(m.Orders, order)
}

func (m *MultipleOrder) AddWithData(clientId, symbol, status string) {
	m.Orders = append(m.Orders, MultipleOrderResponse{
		ClientId: clientId,
		Symbol:   symbol,
		Status:   status,
	})
}

type MultipleOrderResponse struct {
	ClientId string `json:"client_id"`
	Symbol   string `json:"symbol"`
	Status   string `json:"status"`
}
