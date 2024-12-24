package bnfuture

import (
	"context"
	"net/http"

	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	bntraderes "tradething/app/bn/bn_future/bnservice_response_model/trade"

	bncaller "github.com/non26/tradepkg/pkg/bn/binance_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/binance_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/binance_response"
)

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *bntradereq.QueryOrderBinanceServiceRequest,
) (*bntraderes.QueryOrderBinanceServiceResponse, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[bntradereq.QueryOrderBinanceServiceRequest](),
		bnresponse.NewBinanceServiceHttpResponse[bntraderes.QueryOrderBinanceServiceResponse](),
		bfes.httpttransport,
		bfes.httpclient,
	)
	res, err := c.CallBinance(
		request,
		bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		bfes.binanceFutureUrl.QueryOrder,
		http.MethodGet,
		bfes.secretkey,
		bfes.apikey,
		bfes.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
