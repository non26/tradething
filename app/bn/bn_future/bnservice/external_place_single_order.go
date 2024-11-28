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

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *bnfuturereq.PlaceSignleOrderBinanceServiceRequest,
) (*bnfutureres.PlaceSignleOrderBinanceServiceResponse, error) {

	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[bnfuturereq.PlaceSignleOrderBinanceServiceRequest](),
		bnresponse.NewBinanceServiceHttpResponse[bnfutureres.PlaceSignleOrderBinanceServiceResponse](),
		bfes.httpttransport,
		bfes.httpclient,
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
