package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
)

func (bfes *binanceFutureExternalService) SetNewLeverage(
	ctx context.Context,
	request *bnserivcemodelreq.SetLeverageBinanceServiceRequest) error {

	_, err := CallBinance[
		bnserivcemodelreq.SetLeverageBinanceServiceRequest,
		// bnservicemodelres.SetLeverageBinanceServiceResponse,
		struct{},
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
