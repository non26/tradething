package kcservice

import (
	"context"
	"net/http"
	model "tradetoolv2/app/kc/app/model/kcservicemodel/spot/order"
	"tradetoolv2/app/kc/kccommon"
)

func (s *spotKcService) PlaceSpotOrderKcService(
	ctx context.Context,
	req *model.PlaceSpotOrderKcServiceRequest,
) (*model.PlaceSpotOrderKcServiceResponse, error) {

	kcServiceRequest := kccommon.NewHttpKcRequest(
		s.serviceName,
		http.MethodPost,
		s.kcConfig.BaseUrl,
		s.kcConfig.PlaceOrderUrl,
		req,
	)
	kcServiceRequest, err := kcServiceRequest.GetKcNewRequest().
		SetKcHeaders(s.apiKey, s.apiVersion, s.passphass, s.secretKey).
		Error()
	if err != nil {
		return nil, err
	}

	kcServiceResponse := kccommon.NewHttpKcResponse[*model.PlaceSpotOrderKcServiceResponse](
		kcServiceRequest,
		s.serviceName,
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
