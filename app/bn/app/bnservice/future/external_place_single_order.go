package bnservice

import (
	"context"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
)

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest,
) (*bnservicemodelres.PlaceSignleOrderBinanceServiceResponse, error) {

	res, err := CallBinance[
		bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest,
		bnservicemodelres.PlaceSignleOrderBinanceServiceResponse,
	](
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
