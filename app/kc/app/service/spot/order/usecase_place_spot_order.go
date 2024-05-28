package service

import (
	"context"
	model "tradetoolv2/app/kc/app/model/handlermodel/spot/order"
)

func (s *spotOrderService) PlaceSpotOrderService(
	ctx context.Context,
	req *model.PlaceSpotOrderHandlerRequest,
) (*model.PlaceSpotOrderHandlerResponse, error) {
	res, err := s.kcservice.PlaceSpotOrderKcService(
		ctx,
		req.ToPlaceSpotOrderKcServiceRequest(),
	)
	if err != nil {
		return nil, err
	}
	_ = res
	return nil, nil
}
