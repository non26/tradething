package bnfuture

type InvalidateBotHandlerResponse struct {
	BotId      string `json:"bot_id"`
	BotOrderId string `json:"bot_order_id"`
	Message    string `json:"message"`
}
