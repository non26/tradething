package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position_history/dto"
	"tradething/app/bn/future/position_history/service"
)

type positionHistoryExternalService struct {
	positionHistoryService service.IService
}

type IPositionHistoryExternalService interface {
	Get(ctx context.Context, clientId string) (*domain.Order, error)
	Insert(ctx context.Context, positionHistory *domain.Order) error
}

func NewPositionHistoryExternalService(positionHistoryService service.IService) IPositionHistoryExternalService {
	return &positionHistoryExternalService{positionHistoryService: positionHistoryService}
}

func (s *positionHistoryExternalService) Get(ctx context.Context, clientId string) (*domain.Order, error) {
	positionHistory, err := s.positionHistoryService.GetPositionHistory(ctx, clientId)
	if err != nil {
		return nil, err
	}
	positionHistoryDto := dto.NewPositionHistoryDto()
	return positionHistoryDto.ToDomain(positionHistory), nil
}

func (s *positionHistoryExternalService) Insert(ctx context.Context, positionHistory *domain.Order) error {
	positionHistoryDto := dto.NewPositionHistoryDto()
	positionHistoryDomain := positionHistoryDto.FromDomain(positionHistory)
	err := s.positionHistoryService.InsertPositionHistory(ctx, positionHistoryDomain)
	if err != nil {
		return err
	}
	return nil
}
