package bnfuture

import (
	"context"
	"fmt"
	"net/http"
	mkreq "tradething/app/bn/bn_future/bnservice_request_model/market_data"
	mkres "tradething/app/bn/bn_future/bnservice_response_model/market_data"

	bncaller "github.com/non26/tradepkg/pkg/bn/binance_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/binance_response"
)

func (b *bnMarketDataService) GetCandleStickData(ctx context.Context, request *mkreq.CandleStickRequest) (*mkres.CandleStickData, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[mkreq.CandleStickRequest](),
		bnresponse.NewBinanceServiceHttpResponse[mkres.CandleStickData](),
		b.httpttransport,
		b.httpclient,
	).NeedSignature(false)

	uri := fmt.Sprintf("%v?symbol=%v&interval=%v&startTime=%v&endTime=%v", b.binanceFutureUrl.BinanceFutureMarketData.CandleStick, request.Symbol, request.Interval, request.StartTime, request.EndTime)

	res, err := c.CallBinance(
		mkreq.NewCandleStickRequest(request),
		b.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		// b.binanceFutureUrl.BinanceFutureMarketData.CandleStick,
		uri,
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
