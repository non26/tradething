package db

import (
	"context"
	"tradething/app/bn/future/position_history/domain"
	"tradething/app/bn/future/position_history/infrastructure/db/models"
)

type IBnFtHistoryRepository interface {
	Get(ctx context.Context, clientId string) (*models.BnFtHistory, error)
	Insert(ctx context.Context, history *domain.BnFtHistory) error
}
