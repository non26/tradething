package service

import (
	"context"
	bnfuture "tradething/app/bn/bn_future/bnservice"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
	svcFuture "tradething/app/bn/bn_future/service_model"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *svcFuture.SetLeverageServiceRequest,
	) (*bnfutureres.SetLeverageBinanceHandlerResponse, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *svcFuture.PlaceSignleOrderServiceRequest,
	) (*bnfutureres.PlaceSignleOrderHandlerResponse, error)
	QueryOrder(
		ctx context.Context,
		request *svcFuture.QueryOrderServiceRequest,
	) (*bnfutureres.QueryOrderBinanceHandlerResponse, error)
}

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           bnfuture.IBinanceFutureExternalService
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bnfuture.IBinanceFutureExternalService,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
	}
}
