package service

import (
	"context"
	kcservice "tradething/app/kc/app/kcservice/spot/order"
	model "tradething/app/kc/app/model/handlermodel/spot/order"
)

type ISpotOrderService interface {
	PlaceSpotOrderService(
		ctx context.Context,
		req *model.PlaceSpotOrderHandlerRequest,
	) (*model.PlaceSpotOrderHandlerResponse, error)
}

type spotOrderService struct {
	kcservice kcservice.ISpotKcService
}

func NewSpotOrderService(
	kcservice kcservice.ISpotKcService,
) ISpotOrderService {
	return &spotOrderService{
		kcservice,
	}
}
