package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *currentPositionService) ScanWithClientId(ctx context.Context, clientId string) (*domain.Position, error) {
	currentPosition, err := s.repository.ScanWithClientId(ctx, clientId)
	if err != nil {
		return nil, err
	}

	res := &domain.Position{
		ClientId:     currentPosition.ClientId,
		Symbol:       currentPosition.GetSymbol(),
		PositionSide: currentPosition.PositionSide,
		Side:         currentPosition.Side,
		AmountB:      currentPosition.AmountB,
		AccountId:    currentPosition.GetAccountId(),
	}
	return res, nil
}
