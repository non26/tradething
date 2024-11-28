package service

import (
	"context"
	bnfuture "tradething/app/bn/bn_future/bnservice"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
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
}

type binanceFutureService struct {
	binanceFutureServiceName string
	binanceService           bnfuture.IBinanceFutureExternalService
	repository               bndynamodb.IRepository
}

func NewBinanceFutureService(
	binanceFutureServiceName string,
	binanceService bnfuture.IBinanceFutureExternalService,
	repository bndynamodb.IRepository,
) IBinanceFutureService {
	return &binanceFutureService{
		binanceFutureServiceName,
		binanceService,
		repository,
	}
}
