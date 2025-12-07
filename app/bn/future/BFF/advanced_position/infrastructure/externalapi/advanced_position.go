package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/advanced_position/domain"
	"tradething/app/bn/future/BFF/advanced_position/infrastructure/externalapi/dto"
	"tradething/app/bn/future/position/service"
)

type advancedPositionExternalService struct {
	advancedPositionService service.IAdvancedPositionService
}

type IAdvancedPositionExternalService interface {
	Get(ctx context.Context, clientId string) (*domain.AdvancedPosition, error)
	GetAll(ctx context.Context) ([]*domain.AdvancedPosition, error)
	Upsert(ctx context.Context, position *domain.AdvancedPosition) error
	Delete(ctx context.Context, clientId string) error
}

func NewAdvancedPositionExternalService(advancedPositionService service.IAdvancedPositionService) IAdvancedPositionExternalService {
	return &advancedPositionExternalService{advancedPositionService: advancedPositionService}
}

func (s *advancedPositionExternalService) Get(ctx context.Context, clientId string) (*domain.AdvancedPosition, error) {
	advancedPosition, err := s.advancedPositionService.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	if advancedPosition == nil {
		return nil, nil
	}
	advancedPositionDto := dto.NewAdvancedPositionDto()
	return advancedPositionDto.ToDomain(advancedPosition), nil
}

func (s *advancedPositionExternalService) GetAll(ctx context.Context) ([]*domain.AdvancedPosition, error) {
	advancedPositions, err := s.advancedPositionService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if advancedPositions == nil {
		return nil, nil
	}
	advancedPositionDoamin := make([]*domain.AdvancedPosition, len(advancedPositions))
	for i, advancedPosition := range advancedPositions {
		advancedPositionsDto := dto.NewAdvancedPositionDto()
		advancedPositionDoamin[i] = advancedPositionsDto.ToDomain(advancedPosition)
	}
	return advancedPositionDoamin, nil
}

func (s *advancedPositionExternalService) Upsert(ctx context.Context, position *domain.AdvancedPosition) error {
	advancedPositionDto := dto.NewAdvancedPositionDto()
	req := advancedPositionDto.FromDomain(position)
	err := s.advancedPositionService.Upsert(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *advancedPositionExternalService) Delete(ctx context.Context, clientId string) error {
	err := s.advancedPositionService.Delete(ctx, clientId)
	if err != nil {
		return err
	}
	return nil
}
