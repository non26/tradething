package service

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
)

func (s *service) Upsert(ctx context.Context, req *domain.Position) error {
	openingPosition, err := s.positionService.Get(ctx, req.Symbol, req.AccountId, req.PositionSide)
	if err != nil {
		return err
	}

	if openingPosition == nil {
		err := s.positionService.Upsert(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
