package bnfuture

import (
	"context"
	"net/http"

	bnfuturereq "tradething/app/bn/bn_future/bnservice_request_model"
	bnfutureres "tradething/app/bn/bn_future/bnservice_response_model"
	"tradething/app/bn/bncommon"
)

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *bnfuturereq.QueryOrderBinanceServiceRequest,
) (*bnfutureres.QueryOrderBinanceServiceResponse, error) {
	c := NewCallBinance(
		bncommon.NewBinanceServiceHttpRequest[bnfuturereq.QueryOrderBinanceServiceRequest](),
		bncommon.NewBinanceServiceHttpResponse[bnfutureres.QueryOrderBinanceServiceResponse](),
		bncommon.NewBinanceTransport(&http.Transport{}),
		bncommon.NewBinanceSerivceHttpClient(),
	)
	res, err := c.CallBinance(
		request,
		bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		bfes.binanceFutureUrl.QueryOrder,
		http.MethodGet,
		bfes.secrets.BinanceSecretKey,
		bfes.secrets.BinanceApiKey,
		bfes.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
