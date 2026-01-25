package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/position_mangement/domain"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *service) Get(ctx context.Context, req *domain.Position) (*domain.Position, error) {
	cuurentPosition, err := s.positionRepository.Get(ctx, req.ClientId)
	if err != nil {
		return nil, err
	}
	if cuurentPosition != nil {
		return nil, errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	position, err := s.positionService.Get(ctx, req.Symbol, req.AccountId, req.PositionSide)
	if err != nil {
		return nil, err
	}
	return position, nil
}
