package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response_model"
	model "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) CloseBySymbols(
	ctx context.Context,
	request *model.PositionSide,
) (*handlerres.CloseBySymbols, error) {
	response := handlerres.CloseBySymbols{
		Data: []handlerres.CloseBySymbolsData{},
	}
	for _, order := range request.GetData() {
		_, err := b.binanceService.PlaceSingleOrder(ctx, order.ToBinanceServiceModel(b.sideType.Sell()))
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
