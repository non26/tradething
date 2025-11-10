package service

import (
	"context"
	"tradething/app/bn/future/BFF/current_position/domain"
)

func (s *service) Get(ctx context.Context, req *domain.Position) (*domain.Position, error) {
	position, err := s.positionService.GetPosition().Get(ctx, req.Symbol, req.AccountId, req.PositionSide)
	if err != nil {
		return nil, err
	}
	return &domain.Position{
		ClientId:     position.ClientId,
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		Side:         position.Side,
		AmountB:      position.AmountB,
		AccountId:    position.AccountId,
	}, nil
}
