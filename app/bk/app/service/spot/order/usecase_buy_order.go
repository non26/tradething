package service

import (
	"context"
	model "tradetoolv2/app/bk/app/model/handlermodel"
)

func (o *orderService) BuyOrder(
	ctx context.Context,
	req *model.BuyOrderHandlerRequest) (*model.BuyOrderHandlerResponse, error) {
	buyOrderResponse, err := o.bkService.BuyOrder(
		ctx,
		req.ToBuyOrderBkServiceRequest(),
	)
	if err != nil {
		return nil, err
	}
	_ = buyOrderResponse
	return nil, nil
}
