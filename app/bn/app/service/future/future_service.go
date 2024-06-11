package service

import (
	"context"
	bnservice "tradething/app/bn/app/bnservice/future"
	bnhandlerreq "tradething/app/bn/app/model/handlermodel/future/request"
	bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *bnhandlerreq.SetLeverageHandlerRequest) error
	PlaceSingleOrder(
		ctx context.Context,
		request *bnhandlerreq.PlaceSignleOrderHandlerRequest,
	) (*bnhandlerres.PlaceSignleOrderHandlerResponse, error)
	QueryOrder(
		ctx context.Context,
		request *bnhandlerreq.QueryOrderBinanceHandlerRequest,
	) (*bnhandlerres.QueryOrderBinanceHandlerResponse, error)
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
