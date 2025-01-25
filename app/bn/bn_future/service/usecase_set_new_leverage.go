package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *model.Leverage,
) (*handlerres.SetLeverage, serviceerror.IError) {

	res, err := bfs.binanceService.SetNewLeverage(
		ctx,
		request.ToBinanceServiceSetLeverage(),
	)
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_LEVERAGE_ERROR, err)
	}
	return res.ToHandlerResponse(), nil
}
