package externalservice

import (
	"context"
	"tradething/app/bn/future/BFF/market_data/domain"
)

type IMarketData interface {
	GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error)
	GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error)
}
