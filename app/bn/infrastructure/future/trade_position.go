package infrastructure

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type tradePosition struct {
	longPosition  position.IPosition
	shortPosition position.IPosition
}

func NewTradePosition(
	longPosition position.IPosition,
	shortPosition position.IPosition,
) ITradePosition {
	return &tradePosition{
		longPosition:  longPosition,
		shortPosition: shortPosition,
	}
}

func (t *tradePosition) GetPosition(ctx context.Context, position_side string) position.IPosition {
	if position_side == bnconstant.LONG {
		return t.longPosition
	} else if position_side == bnconstant.SHORT {
		return t.shortPosition
	}
	return nil
}
