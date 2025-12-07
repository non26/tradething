package service

import (
	"context"
	"tradething/app/bn/future/BFF/advanced_position/domain"
)

func (s *advancedPositionService) Get(ctx context.Context, clientId string) (*domain.AdvancedPosition, error) {
	positionExt, err := s.advancedService.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}

	return positionExt, nil
}
