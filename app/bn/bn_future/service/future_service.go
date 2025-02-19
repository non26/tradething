package service

import (
	"context"
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *model.Leverage,
	) (*handlerres.SetLeverage, serviceerror.IError)
	PlaceSingleOrder(
		ctx context.Context,
		request *model.Position,
	) (*handlerres.PlacePosition, serviceerror.IError)
	QueryOrder(
		ctx context.Context,
		request *model.Order,
	) (*handlerres.QueryOrder, serviceerror.IError)
	PlaceMultiOrder(
		ctx context.Context,
		request *model.Positions,
	) (*handlerres.PlaceMultiplePosition, serviceerror.IError)
	CloseByClientIds(
		ctx context.Context,
		request *model.ClientIds,
	) (*handlerres.CloseByClientIds, serviceerror.IError)
	CloseBySymbols(
		ctx context.Context,
		request *model.PositionSide,
	) (*handlerres.CloseBySymbols, serviceerror.IError)
	InvalidatePosition(
		ctx context.Context,
		request *model.ClientIds,
	) (*handlerres.InvalidatePosition, serviceerror.IError)
}

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           bntrade.IBinanceFutureExternalService
	bnMarketDataService      bnmarket.IBnMarketDataService
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtQouteUsdtTable       bndynamodb.IBnFtQouteUSDTRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bntrade.IBinanceFutureExternalService,
	bnMarketDataService bnmarket.IBnMarketDataService,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
		bnMarketDataService,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
	}
}
