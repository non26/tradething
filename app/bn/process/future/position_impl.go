package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
	infrastructure "tradething/app/bn/infrastructure/future"
	domain "tradething/app/bn/process/future/domain"

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

func (f *future) PlaceOrder(ctx context.Context, position domain.Position) (*response.Position, error) {
	bnHistory, err := f.bnFtHistoryTable.Get(ctx, position.GetClientId())
	if err != nil {
		return nil, err
	}
	if bnHistory.IsFound() {
		return nil, errors.New("duplicate client id")
	}

	err = f.infraFuture.PlacePosition(ctx, position.ToInfraPosition())
	if err != nil {
		return nil, err
	}

	return &response.Position{
		ClientId: position.GetClientId(),
		Symbol:   position.GetSymbol(),
	}, nil
}
