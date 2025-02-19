package bnfuture

import (
	"context"

	bntradereq "tradething/app/bn/bn_future/bnservice_request/trade"
	bntraderes "tradething/app/bn/bn_future/bnservice_response/trade"
	"tradething/config"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
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
