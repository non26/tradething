package process

import (
	"context"
	infrastructure "tradething/app/bn/internal/infrastructure/future"
	domain "tradething/app/bn/internal/process/future/domain"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type future struct {
	infraFuture              infrastructure.ITrade
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewFuture(
	infraFuture infrastructure.ITrade,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) IFuture {
	return &future{
		infraFuture,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	}
}

func (f *future) PlaceOrder(ctx context.Context, position domain.Position) error {
	bnHistory, err := f.bnFtHistoryTable.Get(ctx, position.GetClientId())
	if err != nil {
		return err
	}
	if bnHistory.IsFound() {
		return err
	}

	err = f.infraFuture.PlacePosition(ctx, position.ToInfraPosition())
	if err != nil {
		return err
	}

	return nil
}
