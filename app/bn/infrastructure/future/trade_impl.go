package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type trade struct {
	tradePosition            ITradePosition
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewTrade(
	tradePosition ITradePosition,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) ITrade {
	return &trade{
		tradePosition:            tradePosition,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (t *trade) PlacePosition(ctx context.Context, position *position.Position) error {
	trade := t.tradePosition.GetPosition(ctx, position.Side)
	if utils.IsBuyPosition(position.Side, position.PositionSide) {
		trade.BuyPosition(ctx, position)
	} else if utils.IsSellPosition(position.Side, position.PositionSide) {
		trade.SellPosition(ctx, position)
	}
	return nil
}
