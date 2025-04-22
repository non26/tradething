package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

func (f *future) SetAdvancedPosition(ctx context.Context, positions []domain.Position) (*response.SetAdvancedPositionResponses, error) {
	responses := response.NewSetAdvancedPositionResponses()
	for _, position := range positions {
		response := response.NewSetAdvancedPositionResponse()
		lookUp, err := f.infraAdvancedPositionLookUp.LookUpByClientId(ctx, position.GetClientId())
		if err != nil {
			response.Fail(position.GetClientId())
			responses.Add(response)
			continue
		}
		_ = lookUp

		err = f.bnFtAdvancedPosition.Upsert(ctx, &dynamodbmodel.BnFtAdvancedPosition{
			ClientID:     position.GetClientId(),
			Symbol:       position.GetSymbol(),
			PositionSide: position.GetPositionSide(),
			Side:         position.GetSide(),
			AmountB:      position.GetEntryQuantity(),
		})
		if err != nil {
			response.Fail(position.GetClientId())
			responses.Add(response)
			continue
		}
		response.Success(position.GetClientId())
		responses.Add(response)
	}

	return responses, nil
}
