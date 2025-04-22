package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
)

func (f *future) GetAdvancedPosition(ctx context.Context, clientId string) (*response.GetAdvancedPositionResponse, error) {
	advancedPosition, err := f.bnFtAdvancedPosition.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}

	if !advancedPosition.IsFound() {
		return nil, errors.New("advanced position not found")
	}

	return &response.GetAdvancedPositionResponse{
		Symbol:       advancedPosition.Symbol,
		PositionSide: advancedPosition.PositionSide,
		Side:         advancedPosition.Side,
		AmountB:      advancedPosition.AmountB,
		ClientId:     advancedPosition.ClientID,
	}, nil
}
