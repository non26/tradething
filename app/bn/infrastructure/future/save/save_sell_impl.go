package save

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"
	domainservice "tradething/app/bn/process/future/domain_service"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type saveSellPosition struct {
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewSaveSellPosition(
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) ISavePositionBySide {
	return &saveSellPosition{bnFtOpeningPositionTable, bnFtCryptoTable, bnFtHistoryTable}
}

func (s *saveSellPosition) Save(ctx context.Context, position *position.Position, lookup *domainservice.LookUp) error {
	err := s.bnFtOpeningPositionTable.Delete(ctx, ToOpeningPositionTable(position))
	if err != nil {
		return err
	}

	err = s.bnFtHistoryTable.Insert(ctx, ToHistoryTable(position))
	if err != nil {
		return err
	}
	return nil
}
