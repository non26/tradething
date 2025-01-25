package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

func (bfs *binanceFutureService) QueryOrder(
	ctx context.Context,
	request *model.Order,
) (*handlerres.QueryOrder, serviceerror.IError) {

	queryOrderRes, err := bfs.binanceService.QueryOrder(
		ctx,
		request.ToBinanceServiceQueryOrder(),
	)
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, err)
	}
	return queryOrderRes.ToHandlerResponse(), nil
}
