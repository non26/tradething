package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/advanced_position/domain"
)

func (s *advancedPositionService) Update(ctx context.Context, position *domain.AdvancedPosition) error {

	existingPosition, err := s.advancedService.Get(ctx, position.ClientId)
	if err != nil {
		return err
	}
	if existingPosition == nil || existingPosition.Symbol == "" {
		return errors.New("position not found")
	}

	err = s.advancedService.Upsert(ctx, position)
	if err != nil {
		return err
	}
	return nil
}
