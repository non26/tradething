package service

import (
	"context"
	"tradething/app/bn/future/BFF/advanced_position/domain"
)

func (s *advancedPositionService) GetAll(ctx context.Context) ([]*domain.AdvancedPosition, error) {
	positionsExt, err := s.advancedService.GetAdvancedPosition().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	positions := make([]*domain.AdvancedPosition, len(positionsExt))
	for i, positionExt := range positionsExt {
		positions[i] = &domain.AdvancedPosition{
			ClientId:     positionExt.ClientId,
			Symbol:       positionExt.Symbol,
			PositionSide: positionExt.PositionSide,
			Side:         positionExt.Side,
			AmountB:      positionExt.AmountB,
			AccountId:    positionExt.AccountId,
		}
	}
	return positions, nil
}
