package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"

	domainservice "tradething/app/bn/process/future/domain_service/close_position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func (f *future) ClosePositionByClientIds(ctx context.Context, clientIds []string) (responses *response.CloseByClientIds, err error) {
	responses = &response.CloseByClientIds{}
	for _, clientId := range clientIds {
		lookUp, err := f.infraClosePositionLookUp.ById(ctx, clientId)
		if err != nil {
			responses.AddWithData(err.Error(), "error", "error", "error", clientId)
			continue
		}

		position := createClosePositionById(lookUp, clientId)
		err = f.infraFuture.PlacePosition(ctx, position.ToInfraPosition())
		if err != nil {
			responses.AddWithData(err.Error(), "error", lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), clientId)
			continue
		}

		err = f.infraSavePosition.Save(ctx, position.ToInfraPosition(), lookUp.ToTradeLookUp())
		if err != nil {
			responses.AddWithData(err.Error(), "error", lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), clientId)
			continue
		}

		responses.AddWithData("success", "success", lookUp.OpeningPosition.GetSymbol(), lookUp.OpeningPosition.GetPositionSide(), clientId)
	}

	return responses, nil
}

func createClosePositionById(lookUp *domainservice.ClsoePositionLookUp, clientId string) *domain.Position {
	position := domain.Position{}
	position.SetClientId(clientId)
	position.SetSymbol(lookUp.OpeningPosition.GetSymbol())
	position.SetPositionSide(lookUp.OpeningPosition.GetPositionSide())
	if lookUp.OpeningPosition.GetPositionSide() == bnconstant.LONG {
		position.SetSide(bnconstant.SELL)
	} else {
		position.SetSide(bnconstant.BUY)
	}
	position.SetEntryQuantity(lookUp.OpeningPosition.GetAmountB())
	return &position
}
