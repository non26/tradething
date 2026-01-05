package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/externalapi/accumulation/dto"
	"tradething/app/bn/future/accumulation/service"
)

type accumulationExternalService struct {
	accumulationService service.IService
}

type IAccumulationExternalService interface {
	Get(ctx context.Context, clientId string) (*domain.Order, error)
	GetAll(ctx context.Context) ([]*domain.Order, error)
	Upsert(ctx context.Context, accumulation *domain.Order) error
	Delete(ctx context.Context, clientId string) error
}

func NewAccumulationExternalService(accumulationService service.IService) IAccumulationExternalService {
	return &accumulationExternalService{accumulationService: accumulationService}
}

func (s *accumulationExternalService) Get(ctx context.Context, clientId string) (*domain.Order, error) {
	accumulation, err := s.accumulationService.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}

	accumulationDto := dto.NewAccumulationDto()
	return accumulationDto.ToDomain(accumulation), nil
}

func (s *accumulationExternalService) GetAll(ctx context.Context) ([]*domain.Order, error) {
	accumulations, err := s.accumulationService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	accumulationDto := dto.NewAccumulationDto()
	accumulationDomains := make([]*domain.Order, len(accumulations))
	for i, accumulation := range accumulations {
		accumulationDomains[i] = accumulationDto.ToDomain(accumulation)
	}
	return accumulationDomains, nil
}

func (s *accumulationExternalService) Upsert(ctx context.Context, accumulation *domain.Order) error {
	accumulationDto := dto.NewAccumulationDto()
	accumulationDomain := accumulationDto.FromDomain(accumulation)
	err := s.accumulationService.Upsert(ctx, accumulationDomain)
	if err != nil {
		return err
	}
	return nil
}

func (s *accumulationExternalService) Delete(ctx context.Context, clientId string) error {
	err := s.accumulationService.Delete(ctx, clientId)
	if err != nil {
		return err
	}
	return nil
}
