package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type tradeBuilder struct {
	longPosition  position.IPosition
	shortPosition position.IPosition
}

func NewTradePosition(
	longPosition position.IPosition,
	shortPosition position.IPosition,
) ITradeBuilder {
	return &tradeBuilder{
		longPosition:  longPosition,
		shortPosition: shortPosition,
	}
}

func (t *tradeBuilder) GetTradePosition(ctx context.Context, position_side string) position.IPosition {
	switch position_side {
	case bnconstant.LONG:
		return t.longPosition
	case bnconstant.SHORT:
		return t.shortPosition
	default:
		return nil
	}
}
