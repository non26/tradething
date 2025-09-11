package infrastructure

import (
	"context"

	"github.com/non26/tradepkg/pkg/bn/utils"

	infra "tradething/app/bn/infrastructure/future"
	builder "tradething/app/bn/infrastructure/future/builder"
	position "tradething/app/bn/infrastructure/future/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type trade struct {
	tradeBuilder             builder.ITradeBuilder
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewTrade(
	tradeBuilder builder.ITradeBuilder,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) infra.ITrade {
	return &trade{
		tradeBuilder:             tradeBuilder,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (t *trade) PlacePosition(ctx context.Context, position *position.Position) error {
	var err error
	trade := t.tradeBuilder.GetTradePosition(ctx, position.PositionSide)
	if utils.IsBuyPosition(position.Side, position.PositionSide) {
		err = trade.BuyPosition(ctx, position)
	} else if utils.IsSellPosition(position.Side, position.PositionSide) {
		err = trade.SellPosition(ctx, position)
	}
	return err
}
