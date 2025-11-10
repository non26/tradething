package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *advancedPositionService) Get(ctx context.Context, clientId string) (*domain.Position, error) {
	results, err := s.repository.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return nil, nil
	}

	return &domain.Position{
		ClientId:     results.ClientID,
		Symbol:       results.Symbol,
		PositionSide: results.PositionSide,
		Side:         results.Side,
		AmountB:      results.AmountB,
		AccountId:    results.AccountId,
	}, nil
}
