package bnfuture

import (
	"context"
	bnservice "tradething/app/bn/bn_future/bnservice"
	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"
)

type IBotService interface {
	TimeIntervalSemiBotService(
		ctx context.Context,
		req *bothandlerreq.TradeTimeIntervalBinanceFutureRequest,
	) (*bothandlerres.TradeTimeIntervalBinanceFutureResponse, error)
}

type botService struct {
	bn_service bnservice.IBinanceFutureExternalService
}

func NewBotService(
	bn_service bnservice.IBinanceFutureExternalService,
) IBotService {
	return &botService{
		bn_service,
	}
}
