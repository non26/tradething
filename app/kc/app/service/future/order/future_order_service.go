package service

import (
	"context"
	kcservice "tradetoolv2/app/kc/app/kcservice/future/order"
	model "tradetoolv2/app/kc/app/model/handlermodel/future/order"
)

type IFutureOrderService interface {
	PlaceFutureOrderService(
		ctx context.Context,
		req *model.PlaceFutureOrderHandlerRequest,
	) (*model.PlaceFutureOrderHandlerResponse, error)
}

type futureOrderService struct {
	kcservice kcservice.IFutureKcService
}

func NewFutureOrderService(
	kcservice kcservice.IFutureKcService,
) IFutureOrderService {
	return &futureOrderService{
		kcservice,
	}
}
