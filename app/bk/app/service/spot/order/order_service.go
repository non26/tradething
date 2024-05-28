package service

import (
	"context"
	bkservice "tradetoolv2/app/bk/app/bkservice/order"
	model "tradetoolv2/app/bk/app/model/handlermodel"
)

type IOrderService interface {
	SellOrder(
		ctx context.Context,
		req *model.SellHandlerRequest) (*model.SellOrderHandlerResponse, error)
	BuyOrder(
		ctx context.Context,
		req *model.BuyOrderHandlerRequest) (*model.BuyOrderHandlerResponse, error)
}

type orderService struct {
	bkService bkservice.IOrderBkService
}

func NewOrderService(
	bkService bkservice.IOrderBkService,
) IOrderService {
	return &orderService{
		bkService: bkService,
	}
}
