package service

import (
	"context"
	"tradething/app/bn/future/position_history/domain"
)

func (s *service) InsertPositionHistory(ctx context.Context, history *domain.BnFtHistory) error {
	err := s.repository.Insert(ctx, history)
	if err != nil {
		return err
	}
	return nil
}
