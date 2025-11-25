package service

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
	externalapi "tradething/app/bn/future/BFF/position_mangement/infrastructure/external_api"
)

type IService interface {
	Get(ctx context.Context, req *domain.Position) (*domain.Position, error)
	GetAll(ctx context.Context) ([]*domain.Position, error)
	Upsert(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, position *domain.Position) error
}

type service struct {
	positionService externalapi.IPositionExternalService
}

func NewService(positionService externalapi.IPositionExternalService) IService {
	return &service{positionService: positionService}
}
