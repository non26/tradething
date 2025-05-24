package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
)

func (f *future) GetAdvancedPosition(ctx context.Context, clientId string) (*response.GetAdvancedPositionResponse, error) {
	lookUp, err := f.infraAdvancedPositionLookUp.LookUpByClientId(ctx, clientId)
	if err != nil {
		return nil, err
	}
	if !lookUp.AdvancedPosition.IsFound() {
		return &response.GetAdvancedPositionResponse{
			ClientId:    clientId,
			FailMessage: "advanced position not found",
		}, nil
	}

	return &response.GetAdvancedPositionResponse{
		Symbol:       lookUp.AdvancedPosition.GetSymbol(),
		PositionSide: lookUp.AdvancedPosition.GetPositionSide(),
		Side:         lookUp.AdvancedPosition.GetSide(),
		AmountB:      lookUp.AdvancedPosition.GetAmountB(),
		ClientId:     lookUp.AdvancedPosition.GetClientId(),
	}, nil
}
