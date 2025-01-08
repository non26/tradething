package service

import (
	"context"
	"errors"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) SetAdvancedPosition(
	ctx context.Context,
	request *svcfuture.SetAdvancedPositionServiceRequest,
) (*svchandlerres.SetAdvancedPositionHandlerResponse, error) {

	dbHistory, err := b.repository.GetHistoryByClientID(ctx, request.GetClientOrderId())
	if err != nil {
		return nil, err
	}
	if dbHistory.IsFound() {
		return nil, errors.New("history not found")
	}

	dbOpenOrder, err := b.repository.GetOpenOrderBySymbolAndPositionSide(ctx, &dynamodbrepository.BnFtOpeningPosition{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
	})
	if err != nil {
		return nil, err
	}
	if dbOpenOrder.IsFound() {
		return nil, errors.New("open order already exists")
	}

	dbAdvancedPosition, err := b.repository.GetAdvancedPositionBySymbolAndPositionSide(ctx, &dynamodbrepository.BnFtAdvancedPositionModel{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
	})
	if err != nil {
		return nil, err
	}
	if dbAdvancedPosition.IsFound() {
		return nil, errors.New("advanced position already exists")
	}

	err = b.repository.InsertAdvancedPosition(ctx, &dynamodbrepository.BnFtAdvancedPositionModel{
		Symbol:       request.GetSymbol(),
		PositionSide: request.GetPositionSide(),
		Side:         request.GetSide(),
		AmountQ:      request.GetEntryQuantity(),
		ClientId:     request.GetClientOrderId(),
	})
	if err != nil {
		return nil, err
	}

	return &svchandlerres.SetAdvancedPositionHandlerResponse{
		PlaceSignleOrderHandlerResponse: svchandlerres.PlaceSignleOrderHandlerResponse{
			Symbol:   request.GetSymbol(),
			Quantity: request.GetEntryQuantity(),
		},
	}, nil
}
