package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *currentPositionService) Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Position, error) {
	results, err := s.repository.Get(ctx, symbol, accountId, positionSide)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return nil, nil
	}
	return &domain.Position{
		ClientId:     results.SymbolAccountId,
		PositionSide: results.PositionSide,
		Side:         results.Side,
		AmountB:      results.AmountB,
		CreatedAt:    results.CreatedAt,
		AccountId:    results.GetAccountId(),
	}, nil
}
