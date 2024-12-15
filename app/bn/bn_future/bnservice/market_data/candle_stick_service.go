package bnfuture

import (
	"context"
	"net/http"
	mkreq "tradething/app/bn/bn_future/bnservice_request_model/market_data"
	mkres "tradething/app/bn/bn_future/bnservice_response_model/market_data"

	bncaller "github.com/non26/tradepkg/pkg/bn/binance_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/binance_response"
)

func (b *bnMarketDataService) GetCandleStickData(ctx context.Context, request *mkreq.CandleStickRequest) (*mkres.CandleStickResponse, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[mkreq.CandleStickRequest](),
		bnresponse.NewBinanceServiceHttpResponse[mkres.CandleStickResponse](),
		b.httpttransport,
		b.httpclient,
	).NeedSignature(false)

	res, err := c.CallBinance(
		mkreq.NewCandleStickRequest(request),
		b.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		b.binanceFutureUrl.BinanceFutureMarketData.CandleStick,
		http.MethodGet,
		b.secrets.BinanceSecretKey,
		b.secrets.BinanceApiKey,
		b.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
