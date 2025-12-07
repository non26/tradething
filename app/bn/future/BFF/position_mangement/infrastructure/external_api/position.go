package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
	"tradething/app/bn/future/BFF/position_mangement/infrastructure/external_api/dto"
	"tradething/app/bn/future/position/service"
)

type positionExternalService struct {
	currentPositionService service.ICurrentPositionService
}

type IPositionExternalService interface {
	Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Position, error)
	GetAll(ctx context.Context) ([]*domain.Position, error)
	Upsert(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, symbol string, accountId string, positionSide string) error
}

func NewPositionService(currentPositionService service.ICurrentPositionService) IPositionExternalService {
	return &positionExternalService{currentPositionService: currentPositionService}
}

func (s *positionExternalService) Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Position, error) {
	currentPosition, err := s.currentPositionService.Get(ctx, symbol, accountId, positionSide)
	if err != nil {
		return nil, err
	}
	currentPositionDto := dto.NewCurrentPositionDto()
	return currentPositionDto.ToDomain(currentPosition), nil
}

func (s *positionExternalService) GetAll(ctx context.Context) ([]*domain.Position, error) {
	currentPositions, err := s.currentPositionService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	currentPositionDomain := make([]*domain.Position, len(currentPositions))

	for i, currentPosition := range currentPositions {
		currentPositionsDto := dto.NewCurrentPositionDto()
		currentPositionDomain[i] = currentPositionsDto.ToDomain(currentPosition)
	}
	return currentPositionDomain, nil
}

func (s *positionExternalService) Upsert(ctx context.Context, position *domain.Position) error {
	currentPositionDto := dto.NewCurrentPositionDto()
	req := currentPositionDto.FromDomain(position)
	err := s.currentPositionService.Upsert(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *positionExternalService) Delete(ctx context.Context, symbol string, accountId string, positionSide string) error {
	err := s.currentPositionService.Delete(ctx, symbol, accountId, positionSide)
	if err != nil {
		return err
	}
	return nil
}
