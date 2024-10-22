package bnfuture

import (
	"context"
	"net/http"

	bnfuturereq "tradething/app/bn/bn_future/bnservice_request_model"
	bnfutureres "tradething/app/bn/bn_future/bnservice_response_model"
	"tradething/app/bn/bncommon"
)

func (bfes *binanceFutureExternalService) SetNewLeverage(
	ctx context.Context,
	request *bnfuturereq.SetLeverageBinanceServiceRequest) (*bnfutureres.SetLeverageBinanceServiceResponse, error) {
	c := NewCallBinance(
		bncommon.NewBinanceServiceHttpRequest[bnfuturereq.SetLeverageBinanceServiceRequest](),
		bncommon.NewBinanceServiceHttpResponse[bnfutureres.SetLeverageBinanceServiceResponse](),
		bfes.httpttransport,
		bfes.httpclient,
	)
	res, err := c.CallBinance(
		request,
		bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		bfes.binanceFutureUrl.SetLeverage,
		http.MethodPost,
		bfes.secrets.BinanceSecretKey,
		bfes.secrets.BinanceApiKey,
		bfes.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
