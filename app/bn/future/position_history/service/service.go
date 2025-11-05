package service

import (
	"context"
	"tradething/app/bn/future/position_history/domain"
	"tradething/app/bn/future/position_history/infrastructure/db"
)

type IService interface {
	InsertPositionHistory(ctx context.Context, history *domain.BnFtHistory) error
	GetPositionHistory(ctx context.Context, clientId string) (*domain.BnFtHistory, error)
}

type service struct {
	repository db.IBnFtHistoryRepository
}

func NewService(repository db.IBnFtHistoryRepository) IService {
	return &service{repository: repository}
}
