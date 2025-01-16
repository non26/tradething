package bnfuture

import (
	"context"

	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	bntraderes "tradething/app/bn/bn_future/bnservice_response_model/trade"
	"tradething/config"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
)

type IBinanceFutureExternalService interface {
	SetNewLeverage(
		ctx context.Context,
		request *bntradereq.SetLeverage,
	) (*bntraderes.SetLeverageData, error)

	PlaceSingleOrder(
		ctx context.Context,
		request *bntradereq.PlacePosition,
	) (*bntraderes.PlacePositionData, error)

	QueryOrder(
		ctx context.Context,
		request *bntradereq.QueryOrder,
	) (*bntraderes.QueryOrderData, error)
}

type binanceFutureExternalService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	apikey                   string
	secretkey                string
	binanceFutureServiceName string
	httpttransport           bntransport.IBinanceServiceHttpTransport
	httpclient               bnclient.IBinanceSerivceHttpClient
}

func NewBinanceFutureExternalService(
	binanceFutureUrl *config.BinanceFutureUrl,
	apikey string,
	secretkey string,
	binanceFutureServiceName string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) IBinanceFutureExternalService {
	return &binanceFutureExternalService{
		binanceFutureUrl,
		apikey,
		secretkey,
		binanceFutureServiceName,
		httpttransport,
		httpclient,
	}
}
