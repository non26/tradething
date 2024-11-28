package bnfuture

import (
	"context"

	bnfuturereq "tradething/app/bn/bn_future/bnservice_request_model"
	bnfutureres "tradething/app/bn/bn_future/bnservice_response_model"
	"tradething/config"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
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
	httpttransport           bntransport.IBinanceServiceHttpTransport
	httpclient               bnclient.IBinanceSerivceHttpClient
}

func NewBinanceFutureExternalService(
	binanceFutureUrl *config.BinanceFutureUrl,
	secrets *config.Secrets,
	binanceFutureServiceName string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) IBinanceFutureExternalService {
	return &binanceFutureExternalService{
		binanceFutureUrl,
		secrets,
		binanceFutureServiceName,
		httpttransport,
		httpclient,
	}
}
