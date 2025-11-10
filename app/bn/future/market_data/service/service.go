package service

import (
	"context"
	"tradething/app/bn/future/market_data/domain"
	"tradething/app/bn/future/market_data/infrastructure/adaptor"
)

type IService interface {
	GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error)
	GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error)
}

type service struct {
	marketDataAdaptor adaptor.IMarketDataAdaptor
}

func NewService(marketDataAdaptor adaptor.IMarketDataAdaptor) IService {
	return &service{marketDataAdaptor: marketDataAdaptor}
}
