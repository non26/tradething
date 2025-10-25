package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"
)

func (f *future) AccumulatePosition(ctx context.Context, accumulatePosition *domain.Position) (*response.Position, error) {

	tradeLookUp, err := f.infraTradeLookUp.LookUp(ctx, accumulatePosition.ToInfraPosition())
	if err != nil {
		return nil, err
	}
	if tradeLookUp.OpeningPosition.IsFound() {
		return nil, errors.New("opening position found")
	}

	exceedMacAccumulatePosition := accumulatePosition.ExceedMaxAccumulatePosition()
	if exceedMacAccumulatePosition {
		return nil, errors.New("exceed max accumulate position")
	}

	bnposition := accumulatePosition.ToInfraPositionWith(tradeLookUp.OpeningPosition.GetClientId(), tradeLookUp.OpeningPosition.GetSymbol(), tradeLookUp.OpeningPosition.GetPositionSide(), tradeLookUp.OpeningPosition.GetSide(), accumulatePosition.GetEntryQuantity())
	err = f.infraTrade.PlacePosition(ctx, bnposition)
	if err != nil {
		return nil, err
	}

	err = f.infraSavePosition.Save(ctx, bnposition, tradeLookUp, nil, nil)
	if err != nil {
		return nil, err
	}

	return &response.Position{
		ClientId: accumulatePosition.GetClientId(),
		Symbol:   accumulatePosition.GetSymbol(),
	}, nil
}
