package service

import (
	"context"
	"tradething/app/bn/future/BFF/market_data/domain"
	externalservice "tradething/app/bn/future/BFF/market_data/infrastructure/external_service"
)

type IService interface {
	GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error)
	GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error)
}

type service struct {
	infra externalservice.IMarketData
}

func NewService(infra externalservice.IMarketData) IService {
	return &service{infra: infra}
}
