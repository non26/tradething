package service

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
)

func (s *service) Get(ctx context.Context, req *domain.Position) (*domain.Position, error) {
	position, err := s.positionService.Get(ctx, req.Symbol, req.AccountId, req.PositionSide)
	if err != nil {
		return nil, err
	}
	return position, nil
}
