package service

import (
	"context"
	"log"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) SetPosition(
	ctx context.Context,
	request *svcfuture.PlaceSignleOrderServiceRequest,
) (*svchandlerres.PlaceSignleOrderHandlerResponse, error) {

	quote, err := b.repository.GetQouteUSDT(ctx, request.GetSymbol())
	if err != nil {
		log.Println("error get qoute usdt", err.Error())
		return nil, err
	}

	isLong := true
	if request.GetPositionSide() == b.positionSideType.Short() {
		isLong = false
	}

	if !quote.IsFound() {
		log.Println("no quote found, insert new quote")
		quote.SetSymbol(request.GetSymbol())
		if isLong {
			quote.SetCountingLong(1)
			quote.SetCountingShort(0)
		} else {
			quote.SetCountingShort(1)
			quote.SetCountingLong(0)
		}

		err = b.repository.InsertNewSymbolQouteUSDT(ctx, quote)
		if err != nil {
			return nil, err
		}
	} else {
		if isLong {
			quote.SetCountingLong(quote.GetCountingLong() + 1)
		} else {
			quote.SetCountingShort(quote.GetCountingShort() + 1)
		}

		err = b.repository.UpdateQouteUSDT(ctx, quote)
		if err != nil {
			return nil, err
		}
	}

	err = b.repository.InsertNewOpenOrder(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	if err != nil {
		log.Println("error new open order", err.Error())
		return nil, err
	}

	return &svchandlerres.PlaceSignleOrderHandlerResponse{
		Symbol:   request.GetSymbol(),
		Quantity: request.GetEntryQuantity(),
	}, nil
}
