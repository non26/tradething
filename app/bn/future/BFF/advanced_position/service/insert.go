package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/advanced_position/domain"
)

func (s *advancedPositionService) Insert(ctx context.Context, position *domain.AdvancedPosition) error {

	existingPosition, err := s.advancedService.Get(ctx, position.ClientId)
	if err != nil {
		return err
	}
	if existingPosition != nil && existingPosition.Symbol != "" {
		return errors.New("position already exists")
	}

	err = s.advancedService.Upsert(ctx, position)
	if err != nil {
		return err
	}
	return nil
}
