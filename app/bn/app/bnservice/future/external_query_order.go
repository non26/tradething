package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnserivcemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/app/bn/bncommon"
)

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *bnserivcemodelreq.QueryOrderBinanceServiceRequest,
) (*bnservicemodelres.QueryOrderBinanceServiceResponse, error) {
	c := NewCallBinance(
		bncommon.NewBinanceServiceHttpRequest[bnserivcemodelreq.QueryOrderBinanceServiceRequest](),
		bncommon.NewBinanceServiceHttpResponse[bnserivcemodelres.QueryOrderBinanceServiceResponse](),
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
