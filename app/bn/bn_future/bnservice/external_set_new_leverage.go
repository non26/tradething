package bnfuture

import (
	"context"
	"net/http"

	bnfuturereq "tradething/app/bn/bn_future/bnservice_request_model"
	bnfutureres "tradething/app/bn/bn_future/bnservice_response_model"

	bncaller "github.com/non26/tradepkg/pkg/bn/binance_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/binance_response"
)

func (bfes *binanceFutureExternalService) SetNewLeverage(
	ctx context.Context,
	request *bnfuturereq.SetLeverageBinanceServiceRequest) (*bnfutureres.SetLeverageBinanceServiceResponse, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[bnfuturereq.SetLeverageBinanceServiceRequest](),
		bnresponse.NewBinanceServiceHttpResponse[bnfutureres.SetLeverageBinanceServiceResponse](),
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
