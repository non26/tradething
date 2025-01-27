package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"

	utils "github.com/non26/tradepkg/pkg/bn/utils"
)

func (b *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *model.Position,
) (*handlerres.PlacePosition, serviceerror.IError) {

	positionHistory, err := b.bnFtHistoryTable.Get(ctx, request.GetClientId())
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, err)
	}
	if positionHistory.IsFound() {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, errors.New("position client id is not valid"))
	}

	openingPositionTable := request.ToBinanceFutureOpeningPositionRepositoryModel()
	dbOpeningPosition, err := b.bnFtOpeningPositionTable.Get(ctx, openingPositionTable)
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
	}

	if !(dbOpeningPosition.IsFound() || utils.IsSellPosition(request.GetPositionSide(), request.GetSide())) {
		// meaing this is new order, no existing order is found
		placeOrderRes, svcerr := b.openPosition(ctx, request)
		if svcerr != nil {
			return nil, svcerr
		}
		return placeOrderRes.ToBnHandlerResponse(), nil
	} else {
		// there is existing order
		if request.IsSellOrder() {
			placeSellOrderRes, svcerr := b.closePosition(ctx, request)
			if svcerr != nil {
				return nil, svcerr
			}
			return placeSellOrderRes.ToBnHandlerResponse(), nil
		}

		if request.GetSymbol() != dbOpeningPosition.Symbol {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("symbol not match"))
		}
		if request.GetPositionSide() != dbOpeningPosition.PositionSide {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("position side not match"))
		}
		if request.GetClientId() == "" {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("client order id is empty"))
		}

		if request.GetClientId() != dbOpeningPosition.ClientId {
			placeOrderRes, svcerr := b.accumulateOrder(ctx, request)
			if svcerr != nil {
				return nil, svcerr
			}
			return placeOrderRes.ToBnHandlerResponse(), nil
		}
	}
	return nil, nil
}
