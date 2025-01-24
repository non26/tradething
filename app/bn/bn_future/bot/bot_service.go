package bnfuture

import (
	"context"
	handlerres "tradething/app/bn/bn_future/bot_handler_response"
	bnbotsvcreq "tradething/app/bn/bn_future/bot_model"

	bntrade "tradething/app/bn/bn_future/bnservice/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
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
}

func NewBotService(
	binanceService bntrade.IBinanceFutureExternalService,
	bnFtBotTable bndynamodb.IBnFtBotRepository,
	bnFtBotOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
) IBotService {
	return &botService{
		binanceService:     binanceService,
		bnFtBotTable:       bnFtBotTable,
		bnFtBotOnRunTable:  bnFtBotOnRunTable,
		bnFtHistoryTable:   bnFtHistoryTable,
		bnFtQouteUsdtTable: bnFtQouteUsdtTable,
	}
}
