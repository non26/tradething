package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *currentPositionService) Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Position, error) {
	currentPosition, err := s.repository.Get(ctx, symbol, accountId, positionSide)
	if err != nil {
		return nil, err
	}
	if currentPosition == nil {
		return nil, nil
	}
	return &domain.Position{
		ClientId:     currentPosition.SymbolAccountId,
		PositionSide: currentPosition.PositionSide,
		Side:         currentPosition.Side,
		AmountB:      currentPosition.AmountB,
		CreatedAt:    currentPosition.CreatedAt,
		AccountId:    currentPosition.GetAccountId(),
		Symbol:       currentPosition.GetSymbol(),
	}, nil
}
