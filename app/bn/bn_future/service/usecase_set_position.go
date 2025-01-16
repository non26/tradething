package service

import (
	"context"
	"errors"
	"log"
	handlerres "tradething/app/bn/bn_future/handler_response_model"
	model "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) SetPosition(
	ctx context.Context,
	request *model.Position,
) (*handlerres.PlacePosition, error) {

	quote, err := b.bnFtQouteUsdtTable.Get(ctx, request.GetSymbol())
	if err != nil {
		return nil, errors.New("get qoute usdt error " + err.Error())
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

		err = b.bnFtQouteUsdtTable.Insert(ctx, quote)
		if err != nil {
			return nil, err
		}
	} else {
		if isLong {
			quote.SetCountingLong(quote.GetCountingLong() + 1)
		} else {
			quote.SetCountingShort(quote.GetCountingShort() + 1)
		}

		err = b.bnFtQouteUsdtTable.Update(ctx, quote)
		if err != nil {
			return nil, errors.New("update quote error " + err.Error())
		}
	}

	err = b.bnFtOpeningPositionTable.Insert(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	if err != nil {
		return nil, errors.New("new open order error " + err.Error())
	}

	return &handlerres.PlacePosition{
		Symbol:   request.GetSymbol(),
		Quantity: request.GetEntryQuantity(),
	}, nil
}
