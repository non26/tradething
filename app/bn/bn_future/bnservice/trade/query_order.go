package bnfuture

import (
	"context"
	"net/http"

	bntradereq "tradething/app/bn/bn_future/bnservice_request/trade"
	bntraderes "tradething/app/bn/bn_future/bnservice_response/trade"

	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
)

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *bntradereq.QueryOrder,
) (*bntraderes.QueryOrderData, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[bntradereq.QueryOrder](),
		bnresponse.NewBinanceServiceHttpResponse[bntraderes.QueryOrderData](),
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
