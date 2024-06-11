package service

import (
	"context"
	bnhandlerreq "tradething/app/bn/app/model/handlermodel/future/request"
	bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"
)

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *bnhandlerreq.PlaceSignleOrderHandlerRequest,
) (*bnhandlerres.PlaceSignleOrderHandlerResponse, error) {

	placeOrderRes, err := bfs.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		return nil, err
	}

	return placeOrderRes.ToBnHandlerResponse(), nil
}
