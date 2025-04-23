package infrastructure

import (
	"context"

	domainservice "tradething/app/bn/process/future/domain_service/advanced_position"

	future "tradething/app/bn/infrastructure/future"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type advancedPositionLookUp struct {
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
	bnFtAdvancedPosition     bndynamodb.IBnFtAdvancedPositionRepository
}

func NewAdvancedPositionLookUp(
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvancedPosition bndynamodb.IBnFtAdvancedPositionRepository,
) future.IAdvancedPositionLookup {
	return &advancedPositionLookUp{
		bnFtOpeningPositionTable,
		bnFtHistoryTable,
		bnFtAdvancedPosition,
	}
}

func (a *advancedPositionLookUp) LookUpByClientId(ctx context.Context, clientId string) (*domainservice.AdvancedPositionLookUp, error) {

	bnAdvancedPosition, err := a.bnFtAdvancedPosition.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}

	look_up := domainservice.NewAdvancedPositionLookUp()
	look_up.AdvancedPosition.SetIsFound(bnAdvancedPosition.IsFound())
	look_up.AdvancedPosition.SetClientId(clientId)
	look_up.AdvancedPosition.SetSymbol(bnAdvancedPosition.Symbol)
	look_up.AdvancedPosition.SetPositionSide(bnAdvancedPosition.PositionSide)
	look_up.AdvancedPosition.SetSide(bnAdvancedPosition.Side)
	look_up.AdvancedPosition.SetAmountB(bnAdvancedPosition.AmountB)
	return look_up, nil
}
