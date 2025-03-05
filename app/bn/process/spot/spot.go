package spot

import (
	"context"
	infrastructure "tradething/app/bn/infrastructure/spot"
	"tradething/app/bn/process/spot/domain"

	res "tradething/app/bn/handlers/spot/res"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type ISpot interface {
	PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error)
	CloseOrderByClientIds(ctx context.Context, clientIds []string) (*res.CloseByClientIds, error)
	MultiplePosition(ctx context.Context, positions []domain.Order) (*res.MultipleOrder, error)
}

type spot struct {
	infraSpotLookUp            infrastructure.ITradeLookUp
	infraSpotSaveOrder         infrastructure.ITradeSaveOrder
	infraClosePositionLookUp   infrastructure.ICloseOrderLookUp
	infraSpot                  infrastructure.ITrade
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewSpot(
	infraSpotLookUp infrastructure.ITradeLookUp,
	infraSpotSaveOrder infrastructure.ITradeSaveOrder,
	infraClosePositionLookUp infrastructure.ICloseOrderLookUp,
	infraSpot infrastructure.ITrade,
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ISpot {
	return &spot{
		infraSpotLookUp:            infraSpotLookUp,
		infraSpotSaveOrder:         infraSpotSaveOrder,
		infraClosePositionLookUp:   infraClosePositionLookUp,
		infraSpot:                  infraSpot,
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}
