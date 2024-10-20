package service

import (
	"context"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
	svcFuture "tradething/app/bn/bn_future/service_model"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *svcFuture.SetLeverageServiceRequest,
) (*bnfutureres.SetLeverageBinanceHandlerResponse, error) {

	res, err := bfs.binanceService.SetNewLeverage(
		ctx,
		request.ToBinanceServiceSetLeverage(),
	)
	if err != nil {
		return nil, err
	}
	return res.ToHandlerResponse(), err
}
