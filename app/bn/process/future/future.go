package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	infrastructure "tradething/app/bn/infrastructure/future"
	domain "tradething/app/bn/process/future/domain"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type IFuture interface {
	PlaceOrder(ctx context.Context, position domain.Position) (*response.Position, error)
	ClosePositionByClientIds(ctx context.Context, clientIds []string) (*response.CloseByClientIds, error)
	MultiplePosition(ctx context.Context, positions []domain.Position) (*response.MultiplePosition, error)
}

type future struct {
	infraFuture              infrastructure.ITrade
	infraLookUp              infrastructure.ITradeLookUp
	infraSavePosition        infrastructure.ISavePosition
	infraClosePositionLookUp infrastructure.IClosePositionLookup
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewFuture(
	infraFuture infrastructure.ITrade,
	infraLookUp infrastructure.ITradeLookUp,
	infraSavePosition infrastructure.ISavePosition,
	infraClosePositionLookUp infrastructure.IClosePositionLookup,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) IFuture {
	return &future{
		infraFuture,
		infraLookUp,
		infraSavePosition,
		infraClosePositionLookUp,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	}
}
