package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position/dto"
	"tradething/app/bn/future/position/service"
)

type advancedPositionExternalService struct {
	advancedPositionService service.IAdvancedPositionService
}

type IAdvancedPositionExternalService interface {
	Get(ctx context.Context, clientId string) (*domain.Order, error)
	GetAll(ctx context.Context) ([]*domain.Order, error)
	Upsert(ctx context.Context, position *domain.Order) error
	Delete(ctx context.Context, clientId string) error
}

func NewAdvancedPositionExternalService(advancedPositionService service.IAdvancedPositionService) IAdvancedPositionExternalService {
	return &advancedPositionExternalService{advancedPositionService: advancedPositionService}
}

func (s *advancedPositionExternalService) Get(ctx context.Context, clientId string) (*domain.Order, error) {
	advancedPosition, err := s.advancedPositionService.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	advancedPositionDto := dto.NewPositionDto()
	return advancedPositionDto.ToDomain(advancedPosition), nil
}

func (s *advancedPositionExternalService) GetAll(ctx context.Context) ([]*domain.Order, error) {
	advancedPositions, err := s.advancedPositionService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	advancedPositionDto := dto.NewPositionDto()
	advancedPositionDomains := make([]*domain.Order, len(advancedPositions))
	for i, advancedPosition := range advancedPositions {
		advancedPositionDomains[i] = advancedPositionDto.ToDomain(advancedPosition)
	}
	return advancedPositionDomains, nil
}

func (s *advancedPositionExternalService) Upsert(ctx context.Context, position *domain.Order) error {
	positionDto := dto.NewPositionDto()
	positionDomain := positionDto.FromDomain(position)
	err := s.advancedPositionService.Upsert(ctx, positionDomain)
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
