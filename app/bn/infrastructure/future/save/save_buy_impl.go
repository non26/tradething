package save

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"

	domainservice "tradething/app/bn/process/future/domain_service"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type saveBuyPosition struct {
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewSaveBuyPosition(
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) ISavePositionBySide {
	return &saveBuyPosition{bnFtOpeningPositionTable, bnFtCryptoTable, bnFtHistoryTable}
}

func (s *saveBuyPosition) Save(ctx context.Context, position *position.Position, lookup *domainservice.LookUp) error {

	if lookup.OpeningPosition.IsFound() {
		err := position.AddMoreAmountB(lookup.OpeningPosition.GetAmountB())
		if err != nil {
			return err
		}
		err = s.bnFtHistoryTable.Insert(ctx, position.ToHistoryTable())
		if err != nil {
			return err
		}
		position.ClientId = lookup.OpeningPosition.GetClientId()
	}
	// this would be Upsert
	err := s.bnFtOpeningPositionTable.Upsert(ctx, ToOpeningPositionTable(position))
	if err != nil {
		return err
	}
	// upsert
	err = s.bnFtCryptoTable.Upsert(ctx, ToCryptoTable(lookup))
	if err != nil {
		return err
	}

	return nil
}
