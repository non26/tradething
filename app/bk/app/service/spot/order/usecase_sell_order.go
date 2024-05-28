package service

import (
	"context"
	model "tradetoolv2/app/bk/app/model/handlermodel"
)

func (o *orderService) SellOrder(
	ctx context.Context,
	req *model.SellHandlerRequest) (*model.SellOrderHandlerResponse, error) {
	sellOrderResponse, err := o.bkService.SellOrder(
		ctx,
		req.ToSellOrderBkRequest(),
	)
	if err != nil {
		return nil, err
	}
	_ = sellOrderResponse
	return nil, nil
}
