package spot

import (
	"context"
	res "tradething/app/bn/handlers/spot/res"
	"tradething/app/bn/process/spot/domain"
)

func (s *spot) MultiplePosition(ctx context.Context, orders []domain.Order) (response *res.MultipleOrder, err error) {
	response = &res.MultipleOrder{}
	for _, order := range orders {
		_, err := s.PlaceOrder(ctx, &order)
		if err != nil {
			response.AddWithData(order.ClientId, order.Symbol, "failed")
			continue
		}
		response.AddWithData(order.ClientId, order.Symbol, "success")
	}
	return response, nil
}
