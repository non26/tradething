package adaptor

import (
	"context"
	"tradething/config"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

type IBinanceFutureMarketService interface {
	GetKlineCandlestick(ctx context.Context) error
}

type binanceFutureMarketAdaptorService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	apikey                   string
	secretkey                string
	binanceFutureServiceName string
	httpttransport           bntransport.IBinanceServiceHttpTransport
	httpclient               bnclient.IBinanceSerivceHttpClient
}

func NewBinanceFutureMarketAdaptorService(
	binanceFutureUrl *config.BinanceFutureUrl,
	apikey string,
	secretkey string,
	binanceFutureServiceName string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) IBinanceFutureMarketService {
	return &binanceFutureMarketAdaptorService{
		binanceFutureUrl:         binanceFutureUrl,
		apikey:                   apikey,
		secretkey:                secretkey,
		binanceFutureServiceName: binanceFutureServiceName,
		httpttransport:           httpttransport,
		httpclient:               httpclient,
	}
}

func (b *binanceFutureMarketAdaptorService) GetKlineCandlestick(ctx context.Context) error {
	return nil
}
