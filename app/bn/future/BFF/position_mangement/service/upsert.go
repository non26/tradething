package service

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
)

func (s *service) Upsert(ctx context.Context, req *domain.Position) error {
	openingPosition, err := s.positionService.GetPosition().Get(ctx, req.Symbol, req.AccountId, req.PositionSide)
	if err != nil {
		return err
	}

	if openingPosition == nil {
		err := s.positionService.GetPosition().Upsert(ctx, req.ToOpeningPositionExtSrv())
		if err != nil {
			return err
		}
	}
	return nil
}
