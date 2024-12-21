package bnfuture

import (
	"context"
	handlerres "tradething/app/bn/bn_future/bot_handler_response_model"
	bnsvcreq "tradething/app/bn/bn_future/bot_service_model"

	bntrade "tradething/app/bn/bn_future/bnservice/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
)

type IBotService interface {
	BotTimeframeExeInterval(ctx context.Context, req *bnsvcreq.BotTimeframeExeIntervalRequest) (*handlerres.BotTimeframeExeIntervalResponse, error)
}

type botService struct {
	binanceService   bntrade.IBinanceFutureExternalService
	repository       bndynamodb.IRepository
	orderType        positionconstant.IOrderType
	positionSideType positionconstant.IPositionSide
	sideType         positionconstant.ISide
}

func NewBotService(
	binanceService bntrade.IBinanceFutureExternalService,
	repository bndynamodb.IRepository,
	orderType positionconstant.IOrderType,
	positionSideType positionconstant.IPositionSide,
	sideType positionconstant.ISide,
) IBotService {
	return &botService{
		binanceService:   binanceService,
		repository:       repository,
		orderType:        orderType,
		positionSideType: positionSideType,
		sideType:         sideType,
	}
}
