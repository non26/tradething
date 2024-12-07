package service

import (
	"context"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

func (s *binanceFutureService) PlaceMultiOrder(
	ctx context.Context,
	request *svcfuture.PlaceMultiOrderServiceRequest,
) (*svchandlerres.PlaceMultipleOrderHandlerResponse, error) {
	return nil, nil
}
