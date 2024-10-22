package bnfuture

import (
	"context"
	bnservice "tradething/app/bn/bn_future/bnservice"
	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"
	bnrepository "tradething/app/bn/bn_future/repository"
	"tradething/app/bn/bncommon"
)

type IBotService interface {
	TimeIntervalSemiBotService(
		ctx context.Context,
		req *bothandlerreq.TradeTimeIntervalBinanceFutureRequest,
	) (*bothandlerres.TradeTimeIntervalBinanceFutureResponse, error)
}

type botService struct {
	bn_service    bnservice.IBinanceFutureExternalService
	bn_repository bnrepository.IRepository
	order_type    bncommon.IOrderType
	position_side bncommon.IPositionSide
	side          bncommon.ISide
}

func NewBotService(
	bn_service bnservice.IBinanceFutureExternalService,
	bn_repository bnrepository.IRepository,
	order_type bncommon.IOrderType,
	position_side bncommon.IPositionSide,
	side bncommon.ISide,
) IBotService {
	return &botService{
		bn_service:    bn_service,
		bn_repository: bn_repository,
		order_type:    order_type,
		position_side: position_side,
		side:          side,
	}
}
