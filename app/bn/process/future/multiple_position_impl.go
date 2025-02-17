package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"
)

func (f *future) MultiplePosition(ctx context.Context, positions []domain.Position) (response *response.MultiplePosition, err error) {
	for _, position := range positions {
		_, err = f.PlaceOrder(ctx, position)
		if err != nil {
			response.AddWithData(position.GetClientId(), position.GetSymbol(), "failed")
			continue
		}

		response.AddWithData(position.GetClientId(), position.GetSymbol(), "success")
	}

	return response, nil
}
