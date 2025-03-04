package save

import (
	"context"
	"tradething/app/bn/infrastructure/spot/order"
	domainservice "tradething/app/bn/process/spot/domain_service/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type saveBuy struct {
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewSaveBuy(
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ISaveOrder {
	return &saveBuy{
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (s *saveBuy) Save(ctx context.Context, order *order.Order, lookup *domainservice.TradeLookUp) error {
	if lookup.OpeningPosition.IsFound() {
		err := order.AddMoreAmountB(order.Quantity)
		if err != nil {
			return err
		}

		err = s.bnSpotHistoryTable.Insert(ctx, order.ToHistoryTable())
		if err != nil {
			return err
		}

		order.NewClientOrderId = lookup.OpeningPosition.GetClientId()
	}

	err := s.bnSpotOpeningPositionTable.Upsert(ctx, order.ToOpeningSpotTable())
	if err != nil {
		return err
	}

	err = s.bnSpotCryptoTable.Update(ctx, ToCrypto(lookup.Symbol.GetSymbol(), lookup.Symbol.GetCounting()))
	if err != nil {
		return err
	}
	return nil
}
