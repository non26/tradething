package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *currentPositionService) GetAll(ctx context.Context) ([]*domain.Position, error) {
	currentPositions, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if currentPositions == nil {
		return []*domain.Position{}, nil
	}

	positions := make([]*domain.Position, len(currentPositions))
	for i, currentPosition := range currentPositions {
		positions[i] = &domain.Position{
			ClientId:     currentPosition.ClientId,
			PositionSide: currentPosition.PositionSide,
			Side:         currentPosition.Side,
			AmountB:      currentPosition.AmountB,
			CreatedAt:    currentPosition.CreatedAt,
			AccountId:    currentPosition.GetAccountId(),
			Symbol:       currentPosition.GetSymbol(),
		}
	}
	return positions, nil
}
