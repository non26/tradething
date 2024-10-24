package bnfuture

import (
	"context"

	bnfuturereq "tradething/app/bn/bn_future/bnservice_request_model"
	bnfutureres "tradething/app/bn/bn_future/bnservice_response_model"
	"tradething/app/bn/bncommon"
	"tradething/config"
)

type IBinanceFutureExternalService interface {
	SetNewLeverage(
		ctx context.Context,
		request *bnfuturereq.SetLeverageBinanceServiceRequest,
	) (*bnfutureres.SetLeverageBinanceServiceResponse, error)

	PlaceSingleOrder(
		ctx context.Context,
		request *bnfuturereq.PlaceSignleOrderBinanceServiceRequest,
	) (*bnfutureres.PlaceSignleOrderBinanceServiceResponse, error)

	QueryOrder(
		ctx context.Context,
		request *bnfuturereq.QueryOrderBinanceServiceRequest,
	) (*bnfutureres.QueryOrderBinanceServiceResponse, error)
}

type binanceFutureExternalService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	secrets                  *config.Secrets
	binanceFutureServiceName string
	httpttransport           bncommon.IBinanceServiceHttpTransport
	httpclient               bncommon.IBinanceSerivceHttpClient
}

func NewBinanceFutureExternalService(
	binanceFutureUrl *config.BinanceFutureUrl,
	secrets *config.Secrets,
	binanceFutureServiceName string,
	httpttransport bncommon.IBinanceServiceHttpTransport,
	httpclient bncommon.IBinanceSerivceHttpClient,
) IBinanceFutureExternalService {
	return &binanceFutureExternalService{
		binanceFutureUrl,
		secrets,
		binanceFutureServiceName,
		httpttransport,
		httpclient,
	}
}
