package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type ITradeBuilder interface {
	GetPosition(ctx context.Context, position_side string) position.IPosition
}

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

func (t *tradeBuilder) GetPosition(ctx context.Context, position_side string) position.IPosition {
	if position_side == bnconstant.LONG {
		return t.longPosition
	} else if position_side == bnconstant.SHORT {
		return t.shortPosition
	}
	return nil
}
