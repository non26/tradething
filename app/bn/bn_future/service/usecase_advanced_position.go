package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response_model"
	model "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) SetAdvancedPosition(
	ctx context.Context,
	request *model.Position,
) (*handlerres.SetAdvancedPosition, error) {

	dbHistory, err := b.bnFtHistoryTable.Get(ctx, request.GetClientOrderId())
	if err != nil {
		return nil, err
	}
	if dbHistory.IsFound() {
		return nil, errors.New("history not found")
	}

	dbOpenOrder, err := b.bnFtOpeningPositionTable.Get(ctx, &dynamodbrepository.BnFtOpeningPosition{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
	})
	if err != nil {
		return nil, err
	}
	if dbOpenOrder.IsFound() {
		return nil, errors.New("open order already exists")
	}

	dbAdvancedPosition, err := b.bnFtAdvancedPositionTable.Get(ctx, &dynamodbrepository.BnFtAdvancedPositionModel{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
	})
	if err != nil {
		return nil, err
	}
	if dbAdvancedPosition.IsFound() {
		return nil, errors.New("advanced position already exists")
	}

	err = b.bnFtAdvancedPositionTable.Insert(ctx, &dynamodbrepository.BnFtAdvancedPositionModel{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
		Side:         request.GetSide(),
		AmountQ:      request.GetEntryQuantity(),
		ClientId:     request.GetClientOrderId(),
	})
	if err != nil {
		return nil, err
	}

	return &handlerres.SetAdvancedPosition{
		PlacePosition: handlerres.PlacePosition{
			Symbol:   request.GetSymbol(),
			Quantity: request.GetEntryQuantity(),
		},
	}, nil
}
