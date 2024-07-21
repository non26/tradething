package bnservice

import (
	"context"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnserivcemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/config"
)

type IBinanceFutureExternalService interface {
	SetNewLeverage(
		ctx context.Context,
		request *bnserivcemodelreq.SetLeverageBinanceServiceRequest,
	) (*bnserivcemodelres.SetLeverageBinanceServiceResponse, error)

	PlaceSingleOrder(
		ctx context.Context,
		request *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest,
	) (*bnservicemodelres.PlaceSignleOrderBinanceServiceResponse, error)

	QueryOrder(
		ctx context.Context,
		request *bnserivcemodelreq.QueryOrderBinanceServiceRequest,
	) (*bnservicemodelres.QueryOrderBinanceServiceResponse, error)
}

type binanceFutureExternalService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	secrets                  *config.Secrets
	binanceFutureServiceName string
}

func NewBinanceFutureExternalService(
	binanceFutureUrl *config.BinanceFutureUrl,
	secrets *config.Secrets,
	binanceFutureServiceName string,
) IBinanceFutureExternalService {
	return &binanceFutureExternalService{
		binanceFutureUrl,
		secrets,
		binanceFutureServiceName,
	}
}
