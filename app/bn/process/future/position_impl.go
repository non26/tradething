package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
	domain "tradething/app/bn/process/future/domain"
)

func (f *future) PlaceOrder(ctx context.Context, position domain.Position) (*response.Position, error) {
	// bnHistory, err := f.bnFtHistoryTable.Get(ctx, position.GetClientId())
	// if err != nil {
	// 	return nil, err
	// }
	// if bnHistory.IsFound() {
	// 	return nil, errors.New("duplicate client id")
	// }

	bnposition := position.ToInfraPosition()
	lookUp, err := f.infraLookUp.LookUp(ctx, bnposition)
	if err != nil {
		return nil, err
	}

	if position.GetClientId() == lookUp.OpeningPosition.GetClientId() {
		return nil, errors.New("duplicate client id")
	}

	err = f.infraFuture.PlacePosition(ctx, bnposition)
	if err != nil {
		return nil, err
	}

	err = f.infraSavePosition.Save(ctx, bnposition, lookUp)
	if err != nil {
		return nil, err
	}

	return &response.Position{
		ClientId: position.GetClientId(),
		Symbol:   position.GetSymbol(),
	}, nil
}
