package bnfuture

import (
	"context"
	handlerres "tradething/app/bn/bn_future/bot_handler_response_model"
	bnbotsvcreq "tradething/app/bn/bn_future/bot_service_model"

	bntrade "tradething/app/bn/bn_future/bnservice/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
)

type IBotService interface {
	// InvalidateBot(ctx context.Context, req *bnbotsvcreq.InvalidateBot) (*handlerres.InvalidateBotHandlerResponse, error)
	BotTimeframeExeInterval(ctx context.Context, req *bnbotsvcreq.BotTimeframeExeIntervalRequest) (*handlerres.BotTimeframeExeIntervalResponse, error)
}

type botService struct {
	binanceService     bntrade.IBinanceFutureExternalService
	bnFtBotTable       bndynamodb.IBnFtBotRepository
	bnFtBotOnRunTable  bndynamodb.IBnFtBotOnRunRepository
	bnFtHistoryTable   bndynamodb.IBnFtHistoryRepository
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository
	orderType          positionconstant.IOrderType
	positionSideType   positionconstant.IPositionSide
	sideType           positionconstant.ISide
}

func NewBotService(
	binanceService bntrade.IBinanceFutureExternalService,
	bnFtBotTable bndynamodb.IBnFtBotRepository,
	bnFtBotOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	orderType positionconstant.IOrderType,
	positionSideType positionconstant.IPositionSide,
	sideType positionconstant.ISide,
) IBotService {
	return &botService{
		binanceService:     binanceService,
		bnFtBotTable:       bnFtBotTable,
		bnFtBotOnRunTable:  bnFtBotOnRunTable,
		bnFtHistoryTable:   bnFtHistoryTable,
		bnFtQouteUsdtTable: bnFtQouteUsdtTable,
		orderType:          orderType,
		positionSideType:   positionSideType,
		sideType:           sideType,
	}
}
