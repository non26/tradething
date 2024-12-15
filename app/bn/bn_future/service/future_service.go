package service

import (
	"context"
	bnfuture "tradething/app/bn/bn_future/bnservice"
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
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
		request *svcfuture.PlaceSignleOrderServiceRequest,
	) (*svchandlerres.PlaceSignleOrderHandlerResponse, error)
	QueryOrder(
		ctx context.Context,
		request *svcfuture.QueryOrderServiceRequest,
	) (*svchandlerres.QueryOrderBinanceHandlerResponse, error)
	PlaceMultiOrder(
		ctx context.Context,
		request *svcfuture.PlaceMultiOrderServiceRequest,
	) (*svchandlerres.PlaceMultipleOrderHandlerResponse, error)
}

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           bnfuture.IBinanceFutureExternalService
	bnMarketDataService      bnmarket.IBnMarketDataService
	repository               bndynamodb.IRepository
	orderType                positionconstant.IOrderType
	positionSideType         positionconstant.IPositionSide
	sideType                 positionconstant.ISide
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bnfuture.IBinanceFutureExternalService,
	bnMarketDataService bnmarket.IBnMarketDataService,
	repository bndynamodb.IRepository,
	orderType positionconstant.IOrderType,
	positionSideType positionconstant.IPositionSide,
	sideType positionconstant.ISide,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
		bnMarketDataService,
		repository,
		orderType,
		positionSideType,
		sideType,
	}
}
