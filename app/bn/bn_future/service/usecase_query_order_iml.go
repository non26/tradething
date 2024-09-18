package service

import (
	"context"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
)

func (bfs *binanceFutureService) QueryOrder(
	ctx context.Context,
	request *bnfuturereq.QueryOrderBinanceHandlerRequest,
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
