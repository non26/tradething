package service

import (
	"context"
	"tradething/app/bn/future/BFF/current_position/domain"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.Position, error) {
	positionsExt, err := s.positionService.GetPosition().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	positions := make([]*domain.Position, len(positionsExt))
	for i, positionExt := range positionsExt {
		positions[i] = &domain.Position{
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
