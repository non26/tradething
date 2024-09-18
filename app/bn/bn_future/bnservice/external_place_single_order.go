package bnfuture

import (
	"context"
	"net/http"
	"tradething/app/bn/bncommon"

	bnfuturereq "tradething/app/bn/bn_future/bnservice_request_model"
	bnfutureres "tradething/app/bn/bn_future/bnservice_response_model"
)

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *bnfuturereq.PlaceSignleOrderBinanceServiceRequest,
) (*bnfutureres.PlaceSignleOrderBinanceServiceResponse, error) {

	c := NewCallBinance(
		bncommon.NewBinanceServiceHttpRequest[bnfuturereq.PlaceSignleOrderBinanceServiceRequest](),
		bncommon.NewBinanceServiceHttpResponse[bnfutureres.PlaceSignleOrderBinanceServiceResponse](),
		bncommon.NewBinanceTransport(&http.Transport{}),
		bncommon.NewBinanceSerivceHttpClient(),
	)

	res, err := c.CallBinance(
		bnfuturereq.NewPlaceSignleOrderBinanceServiceRequest(request),
		bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		bfes.binanceFutureUrl.SingleOrder,
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
