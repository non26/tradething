package service

import (
	"context"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) CloseBySymbols(
	ctx context.Context,
	request *svcfuture.CloseBySymbolsServiceRequest,
) (*svchandlerres.CloseBySymbolsHandlerResponse, error) {
	response := svchandlerres.CloseBySymbolsHandlerResponse{
		Data: []svchandlerres.CloseBySymbolsHandlerResponseData{},
	}
	for _, order := range request.GetData() {
		_, err := b.binanceService.PlaceSingleOrder(ctx, order.ToBinanceServiceModel(b.sideType.Sell()))
		if err != nil {
			response.Data = append(response.Data, svchandlerres.CloseBySymbolsHandlerResponseData{
				Symbol:  order.GetSymbol(),
				Message: err.Error(),
				Status:  "fail",
			})
			continue
		}
		response.Data = append(response.Data, svchandlerres.CloseBySymbolsHandlerResponseData{
			Symbol:  order.GetSymbol(),
			Message: "success",
			Status:  "success",
		})
	}

	return &response, nil
}
