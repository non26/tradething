package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/position_mangement/domain"
	"tradething/app/bn/future/BFF/position_mangement/infrastructure/external_api/dto"
	externalpositionhistoryservice "tradething/app/bn/future/position_history/service"
)

type IHistoryPositionExternalService interface {
	Get(ctx context.Context, clientId string) (*domain.Position, error)
	Insert(ctx context.Context, positionHistory *domain.Position) error
}

type historyPositionExternalService struct {
	historyPositionService externalpositionhistoryservice.IService
}

func NewHistoryPositionExternalService(historyPositionService externalpositionhistoryservice.IService) IHistoryPositionExternalService {
	return &historyPositionExternalService{historyPositionService: historyPositionService}
}

func (s *historyPositionExternalService) Get(ctx context.Context, clientId string) (*domain.Position, error) {
	historyPosition, err := s.historyPositionService.GetPositionHistory(ctx, clientId)
	if err != nil {
		return nil, err
	}
	historyPositionDto := dto.NewBnFtHistoryDto()
	historyPositionDomain := historyPositionDto.FromExternalPositionHistoryDomainToDomain(historyPosition)
	return historyPositionDomain, nil
}

func (s *historyPositionExternalService) Insert(ctx context.Context, positionHistory *domain.Position) error {
	historyPositionDto := dto.NewBnFtHistoryDto()
	historyPositionDomain := historyPositionDto.FromDomainToExternalPositionHistoryDomain(positionHistory)
	err := s.historyPositionService.InsertPositionHistory(ctx, historyPositionDomain)
	if err != nil {
		return err
	}
	return nil
}
