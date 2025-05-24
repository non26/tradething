package res

type GetAdvancedPositionResponse struct {
	ClientId     string `json:"client_id,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	Side         string `json:"side,omitempty"`
	AmountB      string `json:"amount_b,omitempty"`
	PositionSide string `json:"position_side,omitempty"`
	Message      string `json:"message,omitempty"`
}
