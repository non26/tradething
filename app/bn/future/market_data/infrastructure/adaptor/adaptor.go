package adaptor

import (
	"context"
	"tradething/app/bn/future/market_data/domain"
	"tradething/app/bn/future/market_data/infrastructure/adaptor/market_data/res"
)

type IMarketDataAdaptor interface {
	GetKline(ctx context.Context, kline *domain.Kline) (*res.KlineReaderAdaptorResponse, error)
}
