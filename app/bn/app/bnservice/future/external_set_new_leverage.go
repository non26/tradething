package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnserivcemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
)

func (bfes *binanceFutureExternalService) SetNewLeverage(
	ctx context.Context,
	request *bnserivcemodelreq.SetLeverageBinanceServiceRequest) (*bnserivcemodelres.SetLeverageBinanceServiceResponse, error) {

	res, err := CallBinance[
		bnserivcemodelreq.SetLeverageBinanceServiceRequest,
		bnserivcemodelres.SetLeverageBinanceServiceResponse,
	](
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
