package service

import (
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *tradeService) CloseOrderById(ctx context.Context, clientId string) error {
	history, err := s.positionHistoryService.Get(ctx, clientId)
	if err != nil {
		return err
	}
	if history != nil {
		return errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	currentPosition, err := s.currentPositionService.ScanWithClientId(ctx, clientId)
	if err != nil {
		return err
	}
	if currentPosition == nil {
		return errors.New(appresponse.NotFoundOpeningPositionErrorCode)
	}

	currentPosition.ToSellPosition()
	err = s.tradeAdaptor.NewOrder(ctx, currentPosition)
	if err != nil {
		return err
	}

	err = s.currentPositionService.Delete(ctx, currentPosition.Symbol, currentPosition.AccountId, currentPosition.PositionSide)
	if err != nil {
		return err
	}

	err = s.positionHistoryService.Insert(ctx, currentPosition)
	if err != nil {
		return err
	}

	return nil
}
