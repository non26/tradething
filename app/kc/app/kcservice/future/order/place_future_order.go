package kcservice

import (
	"context"
	"net/http"
	model "tradething/app/kc/app/model/kcservicemodel/future/order"
	"tradething/app/kc/kccommon"
)

func (f *futureKcService) PlaceFutureOrderKcService(
	ctx context.Context,
	req *model.PlaceFutureOrderKcServiceRequest,
) (*model.PlaceFutureOrderKcServiceResponse, error) {
	kcServiceRequest := kccommon.NewHttpKcRequest(
		f.serviceName,
		http.MethodPost,
		f.kcConfig.BaseUrl,
		f.kcConfig.PlaceOrderUrl,
		req,
	)
	kcServiceRequest, err := kcServiceRequest.GetKcNewRequest().
		SetKcHeaders(f.apiKey, f.apiVersion, f.passphass, f.secretKey).
		Error()
	if err != nil {
		return nil, err
	}

	kcServiceResponse := kccommon.NewHttpKcResponse[*model.PlaceFutureOrderKcServiceResponse](
		kcServiceRequest,
		f.serviceName,
	)
	kcServiceResponse, err = kcServiceResponse.SetTransport().
		SetClient().
		GetKcResponse().
		GetKcResponseBody().
		Error()
	if err != nil {
		return nil, err
	}

	return kcServiceResponse.GetBody(), nil
}
