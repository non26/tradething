package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position/dto"
	"tradething/app/bn/future/position/service"
)

type currentPositionExternalService struct {
	currentPositionService service.ICurrentPositionService
}
type ICurrentPositionExternalService interface {
	Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Order, error)
	GetAll(ctx context.Context) ([]*domain.Order, error)
	Upsert(ctx context.Context, position *domain.Order) error
	Delete(ctx context.Context, symbol string, accountId string, positionSide string) error
	ScanWithClientId(ctx context.Context, clientId string) (*domain.Order, error)
}

func NewCurrentPositionExternalService(currentPositionService service.ICurrentPositionService) ICurrentPositionExternalService {
	return &currentPositionExternalService{currentPositionService: currentPositionService}
}

func (s *currentPositionExternalService) Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Order, error) {
	currentPosition, err := s.currentPositionService.Get(ctx, symbol, accountId, positionSide)
	if err != nil {
		return nil, err
	}
	currentPositionDto := dto.NewPositionDto()
	return currentPositionDto.ToDomain(currentPosition), nil
}

func (s *currentPositionExternalService) GetAll(ctx context.Context) ([]*domain.Order, error) {
	currentPositions, err := s.currentPositionService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	currentPositionDto := dto.NewPositionDto()
	currentPositionDomains := make([]*domain.Order, len(currentPositions))
	for i, currentPosition := range currentPositions {
		currentPositionDomains[i] = currentPositionDto.ToDomain(currentPosition)
	}
	return currentPositionDomains, nil
}

func (s *currentPositionExternalService) Upsert(ctx context.Context, position *domain.Order) error {
	positionDto := dto.NewPositionDto()
	positionDomain := positionDto.FromDomain(position)
	err := s.currentPositionService.Upsert(ctx, positionDomain)
	if err != nil {
		return err
	}
	return nil
}

func (s *currentPositionExternalService) Delete(ctx context.Context, symbol string, accountId string, positionSide string) error {
	err := s.currentPositionService.Delete(ctx, symbol, accountId, positionSide)
	if err != nil {
		return err
	}
	return nil
}

func (s *currentPositionExternalService) ScanWithClientId(ctx context.Context, clientId string) (*domain.Order, error) {
	currentPosition, err := s.currentPositionService.ScanWithClientId(ctx, clientId)
	if err != nil {
		return nil, err
	}
	currentPositionDto := dto.NewPositionDto()
	return currentPositionDto.ToDomain(currentPosition), nil
}
