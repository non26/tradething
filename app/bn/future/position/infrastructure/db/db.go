package db

import (
	"context"
	"tradething/app/bn/future/position/domain"
	"tradething/app/bn/future/position/infrastructure/db/models"
)

type IBnFtOpeningPositionRepository interface {
	Get(ctx context.Context, symbol string, accountId string, positionSide string) (*models.BnFtOpeningPosition, error)
	GetAll(ctx context.Context) ([]*models.BnFtOpeningPosition, error)
	Upsert(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, symbol string, accountId string, positionSide string) error
}

type IBnFtAdvancedPositionRepository interface {
	Get(ctx context.Context, clientId string) (*models.BnFtAdvancedPosition, error)
	GetAll(ctx context.Context) ([]*models.BnFtAdvancedPosition, error)
	Upsert(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, clientId string) error
}
