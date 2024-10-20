package service

import (
	"context"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
	svcFuture "tradething/app/bn/bn_future/service_model"
)

func (bfs *binanceFutureService) QueryOrder(
	ctx context.Context,
	request *svcFuture.QueryOrderServiceRequest,
) (*bnfutureres.QueryOrderBinanceHandlerResponse, error) {

	queryOrderRes, err := bfs.binanceService.QueryOrder(
		ctx,
		request.ToBinanceServiceQueryOrder(),
	)
	if err != nil {
		return nil, err
	}
	return queryOrderRes.ToHandlerResponse(), nil
}
