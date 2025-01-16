package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
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
		return nil, errors.New("set leverage error " + err.Error())
	}
	return res.ToHandlerResponse(), err
}
