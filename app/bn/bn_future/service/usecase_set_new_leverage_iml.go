package service

import (
	"context"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *bnfuturereq.SetLeverageHandlerRequest,
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
