package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
)

func (bfes *binanceFutureExternalService) SetNewLeverage(
	ctx context.Context,
	request *bnserivcemodelreq.SetLeverageBinanceServiceRequest) error {

	_, err := CallBinance[
		bnserivcemodelreq.SetLeverageBinanceServiceRequest,
		bnservicemodelres.SetLeverageBinanceServiceResponse,
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
		return err
	}

	return nil
}
