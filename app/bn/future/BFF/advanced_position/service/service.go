package service

import (
	"context"
	"tradething/app/bn/future/BFF/advanced_position/domain"
	"tradething/app/bn/future/BFF/advanced_position/infrastructure/externalapi"
)

type advancedPositionService struct {
	advancedService externalapi.IAdvancedPositionExternalService
}

type IAdvancedPositionService interface {
	Get(ctx context.Context, clientId string) (*domain.AdvancedPosition, error)
	GetAll(ctx context.Context) ([]*domain.AdvancedPosition, error)
	Insert(ctx context.Context, position *domain.AdvancedPosition) error
	Update(ctx context.Context, position *domain.AdvancedPosition) error
	Delete(ctx context.Context, clientId string) error
}

func NewAdvancedPositionService(advancedService externalapi.IAdvancedPositionExternalService) IAdvancedPositionService {
	return &advancedPositionService{advancedService: advancedService}
}
