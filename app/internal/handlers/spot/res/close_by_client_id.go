package res

type CloseByClientIds struct {
	Response []CloseByClientIdsResponse
}

func (c *CloseByClientIds) Add(response CloseByClientIdsResponse) {
	c.Response = append(c.Response, response)
}

func (c *CloseByClientIds) AddWithData(clientId, symbol, status string) {
	c.Response = append(c.Response, CloseByClientIdsResponse{
		ClientId: clientId,
		Symbol:   symbol,
		Status:   status,
	})
}

type CloseByClientIdsResponse struct {
	ClientId string `json:"client_id"`
	Symbol   string `json:"symbol"`
	Status   string `json:"status"`
}
