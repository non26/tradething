package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"

	positionconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type trade struct {
	longPosition             position.IPosition
	shortPosition            position.IPosition
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewTrade(
	longPosition position.IPosition,
	shortPosition position.IPosition,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) ITrade {
	return &trade{
		longPosition:             longPosition,
		shortPosition:            shortPosition,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (t *trade) PlacePosition(ctx context.Context, position *position.Position) error {
	if position.IsLongPosition() {
		switch position.Side {
		case positionconstant.BUY:
			err := t.longPosition.BuyPosition(ctx, position)
			if err != nil {
				return err
			}

		case positionconstant.SELL:
			err := t.longPosition.SellPosition(ctx, position)
			if err != nil {
				return err
			}
		}
	} else if position.IsShortPosition() {
		switch position.Side {
		case positionconstant.BUY:
			err := t.shortPosition.SellPosition(ctx, position)
			if err != nil {
				return err
			}
		case positionconstant.SELL:
			err := t.shortPosition.BuyPosition(ctx, position)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
