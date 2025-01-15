package service

import (
	"context"
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *svcfuture.SetLeverageServiceRequest,
	) (*svchandlerres.SetLeverageBinanceHandlerResponse, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *svcfuture.Position,
	) (*svchandlerres.PlaceSignleOrderHandlerResponse, error)
	QueryOrder(
		ctx context.Context,
		request *svcfuture.QueryOrderServiceRequest,
	) (*svchandlerres.QueryOrderBinanceHandlerResponse, error)
	PlaceMultiOrder(
		ctx context.Context,
		request *svcfuture.PlaceMultiOrderServiceRequest,
	) (*svchandlerres.PlaceMultipleOrderHandlerResponse, error)
	CloseByClientIds(
		ctx context.Context,
		request *svcfuture.CloseByClientIdServiceRequest,
	) (*svchandlerres.CloseByClientIdsHandlerResponse, error)
	CloseBySymbols(
		ctx context.Context,
		request *svcfuture.CloseBySymbolsServiceRequest,
	) (*svchandlerres.CloseBySymbolsHandlerResponse, error)
	SetPosition(
		ctx context.Context,
		request *svcfuture.Position,
	) (*svchandlerres.PlaceSignleOrderHandlerResponse, error)
	InvalidatePosition(
		ctx context.Context,
		request *svcfuture.InvalidatePositionServiceRequest,
	) (*svchandlerres.InvalidatePositionHandlerResponse, error)
	SetAdvancedPosition(
		ctx context.Context,
		request *svcfuture.SetAdvancedPositionServiceRequest,
	) (*svchandlerres.SetAdvancedPositionHandlerResponse, error)
}

type binanceFutureService struct {
	binanceFutureServiceName  string
	binanceService            bntrade.IBinanceFutureExternalService
	bnMarketDataService       bnmarket.IBnMarketDataService
	bnFtOpeningPositionTable  bndynamodb.IBnFtOpeningPositionRepository
	bnFtQouteUsdtTable        bndynamodb.IBnFtQouteUSDTRepository
	bnFtHistoryTable          bndynamodb.IBnFtHistoryRepository
	bnFtAdvancedPositionTable bndynamodb.IBnFtAdvancedPositionRepository
	orderType                 positionconstant.IOrderType
	positionSideType          positionconstant.IPositionSide
	sideType                  positionconstant.ISide
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bntrade.IBinanceFutureExternalService,
	bnMarketDataService bnmarket.IBnMarketDataService,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvancedPositionTable bndynamodb.IBnFtAdvancedPositionRepository,
	orderType positionconstant.IOrderType,
	positionSideType positionconstant.IPositionSide,
	sideType positionconstant.ISide,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
		bnMarketDataService,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
		bnFtAdvancedPositionTable,
		orderType,
		positionSideType,
		sideType,
	}
}
