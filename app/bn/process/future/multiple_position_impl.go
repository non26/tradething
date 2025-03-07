package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"
)

func (f *future) MultiplePosition(ctx context.Context, positions []domain.Position) (responses *response.MultiplePosition, err error) {
	responses = &response.MultiplePosition{}
	for _, position := range positions {
		_, err = f.PlaceOrder(ctx, position)
		if err != nil {
			responses.AddWithData(position.GetClientId(), position.GetSymbol(), "failed")
			continue
		}

		responses.AddWithData(position.GetClientId(), position.GetSymbol(), "success")
	}

	return responses, nil
}
