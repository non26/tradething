package service

import (
	"context"
	bnhandlerreq "tradething/app/bn/app/model/handlermodel/future/request"
)

func (bfs *binanceFutureService) SetNewLeverage(
	ctx context.Context,
	request *bnhandlerreq.SetLeverageHandlerRequest) error {

	err := bfs.binanceService.SetNewLeverage(
		ctx,
		request.ToBinanceServiceSetLeverage(),
	)
	if err != nil {
		return err
	}
	return nil
}
