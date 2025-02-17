package infrastructure

import (
	"context"
	"tradething/app/internal/infrastructure/spot/order"
)

type ITrade interface {
	PlaceOrder(ctx context.Context, order order.Order) error
}
