package infrastructure

import (
	"context"
	"tradething/app/internal/infrastructure/spot/order"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type spotTrade struct {
	order order.OrderSpot
}

func NewSpotTrade(
	order order.OrderSpot,
) ITrade {
	return &spotTrade{
		order: order,
	}
}

func (s *spotTrade) PlaceOrder(ctx context.Context, order order.Order) error {
	if order.Side == bnconstant.BUY {
		err := s.order.BuyOrder(ctx, order)
		if err != nil {
			return err
		}
	} else if order.Side == bnconstant.SELL {
		err := s.order.SellOrder(ctx, order)
		if err != nil {
			return err
		}
	}
	return nil
}
