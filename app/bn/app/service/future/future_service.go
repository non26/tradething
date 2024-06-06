package service

import (
	"context"
	bnservice "tradething/app/bn/app/bnservice/future"
	model "tradething/app/bn/app/model/handlermodel/future"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *model.SetLeverageHandlerRequest) error
	PlaceSingleOrder(
		ctx context.Context,
		request *model.PlaceSignleOrderHandlerRequest,
	) (*model.PlaceSignleOrderHandlerResponse, error)
	QueryOrder(
		ctx context.Context,
		request *model.QueryOrderBinanceHandlerRequest,
	) (*model.QueryOrderBinanceHandlerResponse, error)
}

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           bnservice.IBinanceFutureExternalService
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bnservice.IBinanceFutureExternalService,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
	}
}
