package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *advancedPositionService) GetAll(ctx context.Context) ([]*domain.Position, error) {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if results == nil {
		return []*domain.Position{}, nil
	}

	positions := make([]*domain.Position, len(results))
	for i, result := range results {
		positions[i] = &domain.Position{
			ClientId:     result.ClientID,
			Symbol:       result.Symbol,
			PositionSide: result.PositionSide,
			Side:         result.Side,
			AmountB:      result.AmountB,
			AccountId:    result.AccountId,
		}
	}
	return positions, nil
}
