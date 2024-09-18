package service

import (
	"context"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
)

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *bnfuturereq.PlaceSignleOrderHandlerRequest,
) (*bnfutureres.PlaceSignleOrderHandlerResponse, error) {

	placeOrderRes, err := bfs.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		return nil, err
	}

	return placeOrderRes.ToBnHandlerResponse(), nil
}
