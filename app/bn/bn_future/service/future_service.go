package service

import (
	"context"
	bnfuture "tradething/app/bn/bn_future/bnservice"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *bnfuturereq.SetLeverageHandlerRequest,
	) (*bnfutureres.SetLeverageBinanceHandlerResponse, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *bnfuturereq.PlaceSignleOrderHandlerRequest,
	) (*bnfutureres.PlaceSignleOrderHandlerResponse, error)
	QueryOrder(
		ctx context.Context,
		request *bnfuturereq.QueryOrderBinanceHandlerRequest,
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
