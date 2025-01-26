package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

func (b *binanceFutureService) SetAdvancedPosition(
	ctx context.Context,
	request *model.Position,
) (*handlerres.SetAdvancedPosition, serviceerror.IError) {

	dbHistory, err := b.bnFtHistoryTable.Get(ctx, request.GetClientId())
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, err)
	}
	if dbHistory.IsFound() {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, errors.New("history not found"))
	}

	// dbOpenOrder, err := b.bnFtOpeningPositionTable.Get(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	// if err != nil {
	// 	return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
	// }
	// if dbOpenOrder.IsFound() {
	// 	return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("open order already exists"))
	// }

	dbAdvancedPosition, err := b.bnFtAdvancedPositionTable.Get(ctx, request.ToBnAdvancedPositionRepositoryModel())
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_ADVANCED_POSITION_ERROR, err)
	}
	if dbAdvancedPosition.IsFound() {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_ADVANCED_POSITION_ERROR, errors.New("advanced position already exists"))
	}

	err = b.bnFtAdvancedPositionTable.Insert(ctx, request.ToBnAdvancedPositionRepositoryModel())
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_ADVANCED_POSITION_ERROR, err)
	}

	return &handlerres.SetAdvancedPosition{
		PlacePosition: handlerres.PlacePosition{
			Symbol:   request.GetSymbol(),
			Quantity: request.GetEntryQuantity(),
		},
	}, nil
}
