package service

import (
	"context"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

func (s *binanceFutureService) CloseBySymbols(
	ctx context.Context,
	request *svcfuture.CloseBySymbolsServiceRequest,
) (*svchandlerres.CloseBySymbolsHandlerResponse, error) {
	return nil, nil
}
