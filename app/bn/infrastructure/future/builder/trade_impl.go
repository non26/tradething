package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"

	future "tradething/app/bn/infrastructure/future"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type ITradePositionSideBuilder interface {
	GetPosition(ctx context.Context, position_side string) position.IPosition
}

type tradePositionSideBuilder struct {
	longPosition  position.IPosition
	shortPosition position.IPosition
}

func NewTradePosition(
	longPosition position.IPosition,
	shortPosition position.IPosition,
) ITradePositionSideBuilder {
	return &tradePositionSideBuilder{
		longPosition:  longPosition,
		shortPosition: shortPosition,
	}
}

func (t *tradePositionSideBuilder) GetPosition(ctx context.Context, position_side string) position.IPosition {
	if position_side == bnconstant.LONG {
		return t.longPosition
	} else if position_side == bnconstant.SHORT {
		return t.shortPosition
	}
	return nil
}

type trade struct {
	tradePosition            ITradePositionSideBuilder
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewTrade(
	tradePosition ITradePositionSideBuilder,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) future.ITrade {
	return &trade{
		tradePosition:            tradePosition,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (t *trade) PlacePosition(ctx context.Context, position *position.Position) error {
	var err error
	trade := t.tradePosition.GetPosition(ctx, position.PositionSide)
	if utils.IsBuyPosition(position.Side, position.PositionSide) {
		err = trade.BuyPosition(ctx, position)
	} else if utils.IsSellPosition(position.Side, position.PositionSide) {
		err = trade.SellPosition(ctx, position)
	}
	return err
}
