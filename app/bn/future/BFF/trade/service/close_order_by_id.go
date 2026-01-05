package service

import (
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *tradeService) CloseOrderById(ctx context.Context, orderId string) error {
	history, err := s.positionHistoryService.Get(ctx, orderId)
	if err != nil {
		return err
	}
	if history != nil {
		return errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	return nil
}
