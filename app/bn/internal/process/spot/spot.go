package spot

import (
	"context"
	"errors"
	infrastructure "tradething/app/bn/internal/infrastructure/spot"
	"tradething/app/bn/internal/process/spot/domain"

	res "tradething/app/bn/internal/handlers/spot/res"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type ISpot interface {
	PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error)
}

type spot struct {
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
	infraSpot                  infrastructure.ITrade
}

func NewSpot(
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
	infraSpot infrastructure.ITrade,
) ISpot {
	return &spot{
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
		infraSpot:                  infraSpot,
	}
}

func (s *spot) PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error) {

	spotHistory, err := s.bnSpotHistoryTable.Get(ctx, order.ClientId)
	if err != nil {
		return nil, err
	}

	if spotHistory.IsFound() {
		return nil, errors.New("duplicate client id")
	}

	err = s.infraSpot.PlaceOrder(ctx, order.ToInfrastructureOrder())
	if err != nil {
		return nil, err
	}

	return &res.Trade{
		ClientId: order.ClientId,
		Symbol:   order.Symbol,
	}, nil
}
