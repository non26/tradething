package service

import (
	"context"
	bnhandlerreq "tradething/app/bn/app/model/handlermodel/future/request"
	bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *bnhandlerreq.SetLeverageHandlerRequest,
) (*bnhandlerres.SetLeverageBinanceHandlerResponse, error) {

	res, err := bfs.binanceService.SetNewLeverage(
		ctx,
		request.ToBinanceServiceSetLeverage(),
	)
	if err != nil {
		return nil, err
	}
	return res.ToHandlerResponse(), err
}
