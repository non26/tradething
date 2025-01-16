package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) SetAdvancedPosition(
	ctx context.Context,
	request *model.Position,
) (*handlerres.SetAdvancedPosition, error) {

	dbHistory, err := b.bnFtHistoryTable.Get(ctx, request.GetClientOrderId())
	if err != nil {
		return nil, errors.New("get history error " + err.Error())
	}
	if dbHistory.IsFound() {
		return nil, errors.New("history not found")
	}

	dbOpenOrder, err := b.bnFtOpeningPositionTable.Get(ctx, &dynamodbrepository.BnFtOpeningPosition{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
	})
	if err != nil {
		return nil, errors.New("get open order error " + err.Error())
	}
	if dbOpenOrder.IsFound() {
		return nil, errors.New("open order already exists")
	}

	dbAdvancedPosition, err := b.bnFtAdvancedPositionTable.Get(ctx, &dynamodbrepository.BnFtAdvancedPositionModel{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
	})
	if err != nil {
		return nil, errors.New("get advanced position error " + err.Error())
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
		return nil, errors.New("insert advanced position error " + err.Error())
	}

	return &handlerres.SetAdvancedPosition{
		PlacePosition: handlerres.PlacePosition{
			Symbol:   request.GetSymbol(),
			Quantity: request.GetEntryQuantity(),
		},
	}, nil
}
