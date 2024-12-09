package service

import (
	"context"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) PlaceMultiOrder(
	ctx context.Context,
	request *svcfuture.PlaceMultiOrderServiceRequest,
) (*svchandlerres.PlaceMultipleOrderHandlerResponse, error) {
	response := svchandlerres.PlaceMultipleOrderHandlerResponse{}
	for _, order := range request.Orders {
		signleOrderResponse, err := b.PlaceSingleOrder(ctx, &order)
		if err != nil {
			errSignleOrderResponse := svchandlerres.PlaceSignleOrderHandlerResponse{
				Symbol: order.GetSymbol(),
			}
			response.Result = append(response.Result, errSignleOrderResponse)
			continue
		}
		response.Result = append(response.Result, *signleOrderResponse)
	}
	return &response, nil
}
