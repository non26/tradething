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
	) (*handlerres.InvalidatePosition, serviceerror.IError)
	SetAdvancedPosition(
		ctx context.Context,
		request *model.Position,
	) (*handlerres.SetAdvancedPosition, serviceerror.IError)
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
