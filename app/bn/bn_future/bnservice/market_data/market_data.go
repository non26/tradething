package bnfuture

import (
	"context"

	mkreq "tradething/app/bn/bn_future/bnservice_request/market_data"
	mkres "tradething/app/bn/bn_future/bnservice_response/market_data"
	"tradething/config"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

type IBnMarketDataService interface {
	GetCandleStickData(ctx context.Context, request *mkreq.CandleStickRequest) (*mkres.CandleStickData, error)
}

type bnMarketDataService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	secrets                  *config.Secrets
	binanceFutureServiceName string
	httpttransport           bntransport.IBinanceServiceHttpTransport
	httpclient               bnclient.IBinanceSerivceHttpClient
}

func NewBnMarketDataService(
	binanceFutureUrl *config.BinanceFutureUrl,
	secrets *config.Secrets,
	binanceFutureServiceName string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) IBnMarketDataService {
	return &bnMarketDataService{
		binanceFutureUrl:         binanceFutureUrl,
		secrets:                  secrets,
		binanceFutureServiceName: binanceFutureServiceName,
		httpttransport:           httpttransport,
		httpclient:               httpclient,
	}
}
