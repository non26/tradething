package res

type MultiplePosition struct {
	Positions []MultiplePositionResponse
}

func (m *MultiplePosition) Add(position MultiplePositionResponse) {
	m.Positions = append(m.Positions, position)
}

func (m *MultiplePosition) AddWithData(clientId, symbol, status string) {
	m.Positions = append(m.Positions, MultiplePositionResponse{
		ClientId: clientId,
		Symbol:   symbol,
		Status:   status,
	})
}

type MultiplePositionResponse struct {
	ClientId string `json:"clientId"`
	Symbol   string `json:"symbol"`
	Status   string `json:"status"`
}
