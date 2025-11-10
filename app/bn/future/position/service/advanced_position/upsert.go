package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

func (s *advancedPositionService) Upsert(ctx context.Context, position *domain.Position) error {
	err := s.repository.Upsert(ctx, position)
	if err != nil {
		return err
	}
	return nil
}
