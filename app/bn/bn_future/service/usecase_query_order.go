package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"
)

func (bfs *binanceFutureService) QueryOrder(
	ctx context.Context,
	request *model.Order,
) (*handlerres.QueryOrder, error) {

	queryOrderRes, err := bfs.binanceService.QueryOrder(
		ctx,
		request.ToBinanceServiceQueryOrder(),
	)
	if err != nil {
		return nil, errors.New("query order error " + err.Error())
	}
	return queryOrderRes.ToHandlerResponse(), nil
}
