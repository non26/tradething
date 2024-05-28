package service

import (
	"context"
	model "tradetoolv2/app/kc/app/model/handlermodel/future/order"
)

func (f *futureOrderService) PlaceFutureOrderService(
	ctx context.Context,
	req *model.PlaceFutureOrderHandlerRequest,
) (*model.PlaceFutureOrderHandlerResponse, error) {
	res, err := f.kcservice.PlaceFutureOrderKcService(
		ctx,
		req.ToPlaceOrderKcServiceRequest(),
	)
	if err != nil {
		return nil, err
	}
	_ = res
	return nil, nil
}
