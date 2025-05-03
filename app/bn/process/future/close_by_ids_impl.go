package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"
)

func (f *future) ClosePositionByClientIds(ctx context.Context, clientIds []string) (responses *response.CloseByClientIds, err error) {
	responses = &response.CloseByClientIds{}
	for _, clientId := range clientIds {
		position := &domain.Position{}
		position.SetClientId(clientId)
		lookUp, err := f.infraTradeLookUp.LookUp(ctx, position.ToInfraPosition())
		if err != nil {
			responses.AddFailed(err.Error(), "", "", clientId)
			continue
		}
		if !lookUp.OpeningPosition.IsFound() {
			responses.AddFailed("position not found", "", "", clientId)
			continue
		}

		position = position.NewClosePositionFrom(clientId, lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), lookUp.OpeningPosition.GetAmountB())
		err = f.infraTrade.PlacePosition(ctx, position.ToInfraPosition())
		if err != nil {
			responses.AddFailed(err.Error(), lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), clientId)
			continue
		}

		err = f.infraSavePosition.Save(ctx, position.ToInfraPosition(), lookUp, nil, nil)
		if err != nil {
			responses.AddFailed(err.Error(), lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), clientId)
			continue
		}

		responses.AddSuccess(lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), clientId)
	}

	return responses, nil
}
