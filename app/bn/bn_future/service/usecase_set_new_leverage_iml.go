package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response_model"
	model "tradething/app/bn/bn_future/service_model"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *model.Leverage,
) (*handlerres.SetLeverage, error) {

	res, err := bfs.binanceService.SetNewLeverage(
		ctx,
		request.ToBinanceServiceSetLeverage(),
	)
	if err != nil {
		return nil, err
	}
	return res.ToHandlerResponse(), err
}
