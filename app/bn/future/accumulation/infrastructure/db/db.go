package db

import (
	"context"
	"tradething/app/bn/future/accumulation/domain"
	"tradething/app/bn/future/accumulation/infrastructure/db/models"
)

type IBnFtAccumulationRepository interface {
	Get(ctx context.Context, accumID string) (*models.BnFtAccumulation, error)
	GetAll(ctx context.Context) ([]*models.BnFtAccumulation, error)
	Upsert(ctx context.Context, accumulation *domain.Accumulation) error
	Delete(ctx context.Context, accumID string) error
}
