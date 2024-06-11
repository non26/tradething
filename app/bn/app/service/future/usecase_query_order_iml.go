package service

import (
	"context"
	bnhandlerreq "tradething/app/bn/app/model/handlermodel/future/request"
	bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"
)

func (bfs *binanceFutureService) QueryOrder(
	ctx context.Context,
	request *bnhandlerreq.QueryOrderBinanceHandlerRequest,
) (*bnhandlerres.QueryOrderBinanceHandlerResponse, error) {

	queryOrderRes, err := bfs.binanceService.QueryOrder(
		ctx,
		request.ToBinanceServiceQueryOrder(),
	)
	if err != nil {
		return nil, err
	}
	return queryOrderRes.ToHandlerResponse(), nil
}
