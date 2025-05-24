package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"
)

func (f *future) ManagePosition(ctx context.Context, closeClientId []string, advPositionClientId []string) (*response.ManagePositionRes, error) {
	response := &response.ManagePositionRes{}

	closePosition, err := f.ClosePositionByClientIds(ctx, closeClientId)
	if err != nil {
		return nil, err
	}
	response.ClosePosition = closePosition

	for _, clientId := range advPositionClientId {
		advPosition := domain.NewPosition().SetClientId(clientId)
		advPositionRes, err := f.PlaceOrder(ctx, advPosition)
		if err != nil {
			return nil, err
		}
		response.AdvancedPosition = append(response.AdvancedPosition, advPositionRes)
	}

	return response, nil
}
