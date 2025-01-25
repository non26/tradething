package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

func (b *binanceFutureService) CloseBySymbols(
	ctx context.Context,
	request *model.PositionSide,
) (*handlerres.CloseBySymbols, serviceerror.IError) {
	response := handlerres.CloseBySymbols{
		Data: []handlerres.CloseBySymbolsData{},
	}
	for _, order := range request.GetData() {
		_, err := b.binanceService.PlaceSingleOrder(ctx, order.ToBinanceServiceModel(bnconstant.SELL))
		if err != nil {
			response.Data = append(response.Data, handlerres.CloseBySymbolsData{
				Symbol:  order.GetSymbol(),
				Message: err.Error(),
				Status:  "fail",
			})
			continue
		}
		response.Data = append(response.Data, handlerres.CloseBySymbolsData{
			Symbol:  order.GetSymbol(),
			Message: "success",
			Status:  "success",
		})
	}

	return &response, nil
}
