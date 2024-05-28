package service

import (
	"context"
	bnservice "tradetoolv2/app/bn/app/bnservice/future"
	model "tradetoolv2/app/bn/app/model/handlermodel/future"
)

type IBinanceFutureService interface {
	SetNewLeverage(
		ctx context.Context,
		request *model.SetLeverageHandlerRequest) error
	PlaceSingleOrder(
		ctx context.Context,
		request *model.PlaceSignleOrderHandlerRequest,
	) (*model.PlaceSignleOrderHandlerResponse, error)
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
