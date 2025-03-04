package save

import (
	"context"
	"tradething/app/bn/infrastructure/spot/order"
	domainservice "tradething/app/bn/process/spot/domain_service/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type saveSell struct {
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewSaveSell(
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ISaveOrder {
	return &saveSell{
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (s *saveSell) Save(ctx context.Context, order *order.Order, lookup *domainservice.TradeLookUp) error {
	err := s.bnSpotOpeningPositionTable.Delete(ctx, order.ToOpeningSpotTable())
	if err != nil {
		return err
	}

	err = s.bnSpotHistoryTable.Insert(ctx, order.ToHistoryTable())
	if err != nil {
		return err
	}
	return nil
}
