package service

import (
	"context"
	handleres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

func (b *binanceFutureService) PlaceMultiOrder(
	ctx context.Context,
	request *model.Positions,
) (*handleres.PlaceMultiplePosition, serviceerror.IError) {
	response := handleres.PlaceMultiplePosition{}
	for _, order := range request.Positions {
		signleOrderResponse, svcerr := b.PlaceSingleOrder(ctx, &order)
		if svcerr != nil {
			response.Result.Failed = append(response.Result.Failed, handleres.PlaceMultiplePositionFailed{
				Symbol: order.GetSymbol(),
				Error:  svcerr.Error(),
			})
			continue
		}
		response.Result.Success = append(response.Result.Success, signleOrderResponse.Symbol)
	}
	return &response, nil
}
