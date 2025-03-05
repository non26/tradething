package spot

import (
	"context"
	res "tradething/app/bn/handlers/spot/res"
	"tradething/app/bn/process/spot/domain"
)

func (s *spot) PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error) {

	lookup, err := s.infraSpotLookUp.LookUp(ctx, order.ToInfrastructureOrder())
	if err != nil {
		return nil, err
	}

	err = s.infraSpot.PlaceOrder(ctx, order.ToInfrastructureOrder())
	if err != nil {
		return nil, err
	}

	err = s.infraSpotSaveOrder.Save(ctx, order.ToInfrastructureOrder(), lookup)
	if err != nil {
		return nil, err
	}

	return &res.Trade{
		ClientId: order.ClientId,
		Symbol:   order.Symbol,
	}, nil
}
