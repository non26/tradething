package service

import (
	"context"
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *model.Leverage,
	) (*handlerres.SetLeverage, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *model.Position,
	) (*handlerres.PlacePosition, error)
	QueryOrder(
		ctx context.Context,
		request *model.Order,
	) (*handlerres.QueryOrder, error)
	PlaceMultiOrder(
		ctx context.Context,
		request *model.Positions,
	) (*handlerres.PlaceMultiplePosition, error)
	CloseByClientIds(
		ctx context.Context,
		request *model.ClientIds,
	) (*handlerres.CloseByClientIds, error)
	CloseBySymbols(
		ctx context.Context,
		request *model.PositionSide,
	) (*handlerres.CloseBySymbols, error)
	SetPosition(
		ctx context.Context,
		request *model.Position,
	) (*handlerres.PlacePosition, error)
	InvalidatePosition(
		ctx context.Context,
		request *model.ClientIds,
	) (*handlerres.InvalidatePosition, error)
	SetAdvancedPosition(
		ctx context.Context,
		request *model.Position,
	) (*handlerres.SetAdvancedPosition, error)
}

type binanceFutureService struct {
	binanceFutureServiceName  string
	binanceService            bntrade.IBinanceFutureExternalService
	bnMarketDataService       bnmarket.IBnMarketDataService
	bnFtOpeningPositionTable  bndynamodb.IBnFtOpeningPositionRepository
	bnFtQouteUsdtTable        bndynamodb.IBnFtQouteUSDTRepository
	bnFtHistoryTable          bndynamodb.IBnFtHistoryRepository
	bnFtAdvancedPositionTable bndynamodb.IBnFtAdvancedPositionRepository
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bntrade.IBinanceFutureExternalService,
	bnMarketDataService bnmarket.IBnMarketDataService,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvancedPositionTable bndynamodb.IBnFtAdvancedPositionRepository,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
		bnMarketDataService,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
		bnFtAdvancedPositionTable,
	}
}
