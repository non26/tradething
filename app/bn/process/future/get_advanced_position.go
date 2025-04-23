package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
)

func (f *future) GetAdvancedPosition(ctx context.Context, clientId string) (*response.GetAdvancedPositionResponse, error) {
	lookUp, err := f.infraAdvancedPositionLookUp.LookUpByClientId(ctx, clientId)
	if err != nil {
		return nil, err
	}
	if !lookUp.AdvancedPosition.IsFound() {
		return nil, errors.New("advanced position not found")
	}

	return &response.GetAdvancedPositionResponse{
		Symbol:       lookUp.AdvancedPosition.GetSymbol(),
		PositionSide: lookUp.AdvancedPosition.GetPositionSide(),
		Side:         lookUp.AdvancedPosition.GetSide(),
		AmountB:      lookUp.AdvancedPosition.GetAmountB(),
		ClientId:     lookUp.AdvancedPosition.GetClientId(),
	}, nil
}
