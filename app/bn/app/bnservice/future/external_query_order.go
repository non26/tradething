package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
)

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *bnserivcemodelreq.QueryOrderBinanceServiceRequest,
) (*bnservicemodelres.QueryOrderBinanceServiceResponse, error) {
	res, err := CallBinance[
		bnserivcemodelreq.QueryOrderBinanceServiceRequest,
		bnservicemodelres.QueryOrderBinanceServiceResponse,
	](
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
