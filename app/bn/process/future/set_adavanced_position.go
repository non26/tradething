package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

func (f *future) SetAdvancedPosition(ctx context.Context, positions []*domain.Position) (*response.SetAdvancedPositionResponses, error) {
	responses := response.NewSetAdvancedPositionResponses()
	for _, position := range positions {
		response := response.NewSetAdvancedPositionResponse()
		tradeLookUp, err := f.infraTradeLookUp.LookUp(ctx, position.ToInfraPosition())
		if err != nil {
			response.Fail(position.GetClientId(), err.Error())
			responses.Add(response)
			continue
		}
		if tradeLookUp.OpeningPosition.IsFound() {
			response.Fail(position.GetClientId(), "Opening position found")
			responses.Add(response)
			continue
		}

		AdvLookUp, err := f.infraAdvancedPositionLookUp.LookUpByClientId(ctx, position.GetClientId())
		if err != nil {
			response.Fail(position.GetClientId(), err.Error())
			responses.Add(response)
			continue
		}
		if AdvLookUp.AdvancedPosition.IsFound() {
			response.Fail(position.GetClientId(), "Advanced position found")
			responses.Add(response)
			continue
		}

		advancedPosition := dynamodbmodel.NewBnFtAdvancedPosition()
		advancedPosition.ClientID = position.GetClientId()
		advancedPosition.Symbol = position.GetSymbol()
		advancedPosition.PositionSide = position.GetPositionSide()
		advancedPosition.Side = position.GetSide()
		advancedPosition.AmountB = position.GetEntryQuantity()
		err = f.bnFtAdvancedPosition.Upsert(ctx, advancedPosition)
		if err != nil {
			response.Fail(position.GetClientId(), err.Error())
			responses.Add(response)
			continue
		}
		response.Success(position.GetClientId())
		responses.Add(response)
	}

	return responses, nil
}
