package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/app/bn/bncommon"
)

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest,
) (*bnservicemodelres.PlaceSignleOrderBinanceServiceResponse, error) {

	c := NewCallBinance(
		bncommon.NewBinanceServiceHttpRequest[bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest](),
		bncommon.NewBinanceServiceHttpResponse[bnservicemodelres.PlaceSignleOrderBinanceServiceResponse](),
		bncommon.NewBinanceTransport(&http.Transport{}),
		bncommon.NewBinanceSerivceHttpClient(),
	)

	res, err := c.CallBinance(
		bnserivcemodelreq.NewPlaceSignleOrderBinanceServiceRequest(request),
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
