package service

import (
	"context"
	handleres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) PlaceMultiOrder(
	ctx context.Context,
	request *model.Positions,
) (*handleres.PlaceMultiplePosition, error) {
	response := handleres.PlaceMultiplePosition{}
	for _, order := range request.Positions {
		signleOrderResponse, err := b.PlaceSingleOrder(ctx, &order)
		if err != nil {
			errSignleOrderResponse := handleres.PlacePosition{
				Symbol: order.GetSymbol(),
			}
			response.Result = append(response.Result, errSignleOrderResponse)
			continue
		}
		response.Result = append(response.Result, *signleOrderResponse)
	}
	return &response, nil
}
