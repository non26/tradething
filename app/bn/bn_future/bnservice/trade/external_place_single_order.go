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

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *bntradereq.PlaceSignleOrderBinanceServiceRequest,
) (*bntraderes.PlaceSignleOrderBinanceServiceResponse, error) {

	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[bntradereq.PlaceSignleOrderBinanceServiceRequest](),
		bnresponse.NewBinanceServiceHttpResponse[bntraderes.PlaceSignleOrderBinanceServiceResponse](),
		bfes.httpttransport,
		bfes.httpclient,
	)

	res, err := c.CallBinance(
		bntradereq.NewPlaceSignleOrderBinanceServiceRequest(request),
		bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		bfes.binanceFutureUrl.SingleOrder,
		http.MethodPost,
		bfes.secretkey,
		bfes.apikey,
		bfes.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
