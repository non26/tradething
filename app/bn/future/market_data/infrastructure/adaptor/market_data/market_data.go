package marketdata

import "tradething/app/bn/future/market_data/infrastructure/adaptor"

type marketDataAdaptor struct {
	baseUrl                   string
	klinesCandleStickEndpoint string
}

func NewMarketDataAdaptor(baseUrl string, klinesCandleStickEndpoint string) adaptor.IMarketDataAdaptor {
	return &marketDataAdaptor{baseUrl: baseUrl, klinesCandleStickEndpoint: klinesCandleStickEndpoint}
}
