package service

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.Position, error) {
	positions, err := s.positionService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return positions, nil
}
