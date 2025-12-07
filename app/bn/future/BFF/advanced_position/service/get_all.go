package service

import (
	"context"
	"tradething/app/bn/future/BFF/advanced_position/domain"
)

func (s *advancedPositionService) GetAll(ctx context.Context) ([]*domain.AdvancedPosition, error) {
	positionsExt, err := s.advancedService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return positionsExt, nil
}
