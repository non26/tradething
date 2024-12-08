package service

import (
	"context"
	"errors"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
	svcFuture "tradething/app/bn/bn_future/service_model"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *svcFuture.PlaceSignleOrderServiceRequest,
) (*bnfutureres.PlaceSignleOrderHandlerResponse, error) {

	openingOrder, err := bfs.repository.GetOpenOrderByKey(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel().GetKeyByPositionSideAndSymbol())
	if err != nil {
		return nil, err
	}

	if openingOrder.ClientId == request.GetClientOrderId() {
		return nil, errors.New("opening order already exist")
	}

	placeOrderRes, err := bfs.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		return nil, err
	}

	if openingOrder.IsEmpty() {
		err = bfs.repository.NewOpenOrder(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
		if err != nil {
			return nil, err
		}
	} else {
		if request.GetSide() == openingOrder.Side && request.GetPositionSide() == openingOrder.PositionSide {
			if utils.IsBuyCrypto(request.GetSide(), request.GetPositionSide()) {
				request.AddEntryQuantity(openingOrder.AmountQ)
				err = bfs.repository.UpdateOpenOrder(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
				if err != nil {
					return nil, err
				}
			}
		} else {
			err = bfs.repository.DeleteOpenOrderByKey(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel().GetKeyByPositionSideAndSymbol())
			if err != nil {
				return nil, err
			}
		}
	}

	return placeOrderRes.ToBnHandlerResponse(), nil
}
