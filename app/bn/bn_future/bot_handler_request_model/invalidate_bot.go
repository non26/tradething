package bnfuture

import bnbotsvcreq "tradething/app/bn/bn_future/bot_service_model"

type InvalidateBotHandlerRequest struct {
	BotId      string `json:"bot_id"`
	BotOrderId string `json:"bot_order_id"`
}

func (b *InvalidateBotHandlerRequest) ToServiceModel() *bnbotsvcreq.InvalidateBot {
	return &bnbotsvcreq.InvalidateBot{
		BotId:      b.BotId,
		BotOrderId: b.BotOrderId,
	}
}
