package service

import (
	"context"
	"tradething/app/bn/future/accumulation/domain"
	"tradething/app/bn/future/accumulation/infrastructure/db"
)

type IService interface {
	Get(ctx context.Context, clientId string) (*domain.Accumulation, error)
	GetAll(ctx context.Context) ([]*domain.Accumulation, error)
	Upsert(ctx context.Context, accumulation *domain.Accumulation) error
	Delete(ctx context.Context, clientId string) error
}

type BnFtAccumulationService struct {
	repository db.IBnFtAccumulationRepository
}

func NewBnFtAccumulationService(repository db.IBnFtAccumulationRepository) *BnFtAccumulationService {
	return &BnFtAccumulationService{repository: repository}
}
