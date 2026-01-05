package adaptor

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
)

type ITradeAdaptor interface {
	NewOrder(ctx context.Context, order *domain.Order) error
}
