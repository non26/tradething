package res

type CloseByClientIds struct {
	Response []CloseByClientIdsResponse
}

func (c *CloseByClientIds) Add(response CloseByClientIdsResponse) {
	c.Response = append(c.Response, response)
}

func (c *CloseByClientIds) AddWithData(message string, code string, symbol string, positionSide string, clientId string) {
	c.Response = append(c.Response, CloseByClientIdsResponse{
		Message:      message,
		Code:         code,
		Symbol:       symbol,
		PositionSide: positionSide,
		ClientId:     clientId,
	})
}

type CloseByClientIdsResponse struct {
	Message      string `json:"message"`
	Code         string `json:"code"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"positionSide"`
	ClientId     string `json:"clientId"`
}
