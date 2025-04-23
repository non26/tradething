package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	infra "tradething/app/bn/infrastructure/future"
	infrastructure "tradething/app/bn/infrastructure/future"

	domain "tradething/app/bn/process/future/domain"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type IFuture interface {
	PlaceOrder(ctx context.Context, position domain.Position) (*response.Position, error)
	ClosePositionByClientIds(ctx context.Context, clientIds []string) (*response.CloseByClientIds, error)
	MultiplePosition(ctx context.Context, positions []domain.Position) (*response.MultiplePosition, error)
	SetAdvancedPosition(ctx context.Context, position []domain.Position) (*response.SetAdvancedPositionResponses, error)
	GetAdvancedPosition(ctx context.Context, clientId string) (*response.GetAdvancedPositionResponse, error)
}

type future struct {
	infraTradeBuilder infra.ITrade
	infraTradeLookUp  infra.ITradeLookUp
	infraSavePosition infra.ITradeSavePosition
	// infraClosePositionLookUp    infrastructure.IClosePositionLookup
	infraAdvancedPositionLookUp infra.IAdvancedPositionLookup
	infraCryptoLookUp           infra.ICryptoLookUp
	bnFtOpeningPositionTable    bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable             bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable            bndynamodb.IBnFtHistoryRepository
	bnFtAdvancedPosition        bndynamodb.IBnFtAdvancedPositionRepository
}

func NewFuture(
	infraFuture infrastructure.ITrade,
	infraLookUp infrastructure.ITradeLookUp,
	infraSavePosition infrastructure.ITradeSavePosition,
	// infraClosePositionLookUp infrastructure.IClosePositionLookup,
	infraAdvancedPositionLookUp infrastructure.IAdvancedPositionLookup,
	infraCryptoLookUp infrastructure.ICryptoLookUp,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvancedPosition bndynamodb.IBnFtAdvancedPositionRepository,
) IFuture {
	return &future{
		infraFuture,
		infraLookUp,
		infraSavePosition,
		// infraClosePositionLookUp,
		infraAdvancedPositionLookUp,
		infraCryptoLookUp,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
		bnFtAdvancedPosition,
	}
}
