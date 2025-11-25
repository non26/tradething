package service

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
)

func (s *service) Delete(ctx context.Context, req *domain.Position) error {
	err := s.positionService.GetPosition().Delete(ctx, req.Symbol, req.AccountId, req.PositionSide)
	if err != nil {
		return err
	}
	return nil
}
