package service

import (
	"context"
	"tradething/app/bn/future/position/domain"
)

type ICurrentPositionService interface {
	Get(ctx context.Context, symbol string, accountId string, positionSide string) (*domain.Position, error)
	GetAll(ctx context.Context) ([]*domain.Position, error)
	Upsert(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, symbol string, accountId string, positionSide string) error
	ScanWithClientId(ctx context.Context, clientId string) (*domain.Position, error)
}

type IAdvancedPositionService interface {
	Get(ctx context.Context, clientId string) (*domain.Position, error)
	GetAll(ctx context.Context) ([]*domain.Position, error)
	Upsert(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, clientId string) error
}
