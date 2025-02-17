package order

import (
	"context"
)

type IOrder interface {
	BuyOrder(ctx context.Context, order *Order) error
	SellOrder(ctx context.Context, order *Order) error
}
