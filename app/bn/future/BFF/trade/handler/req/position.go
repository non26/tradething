package req

type Position struct {
	PositionSide string `json:"position_side" binding:"required"`
	Side         string `json:"side" `
	AmountB      string `json:"amount_b"`
	Symbol       string `json:"symbol"`
	ClientId     string `json:"client_id"`
}
