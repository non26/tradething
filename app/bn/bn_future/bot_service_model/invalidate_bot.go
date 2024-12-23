package bnfuture

type InvalidateBot struct {
	BotId        string `json:"bot_id"`
	BotOrderId   string `json:"bot_order_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
}
