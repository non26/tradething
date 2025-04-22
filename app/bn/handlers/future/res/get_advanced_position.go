package res

type GetAdvancedPositionResponse struct {
	ClientId     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	AmountB      string `json:"amount_b"`
	PositionSide string `json:"position_side"`
}
