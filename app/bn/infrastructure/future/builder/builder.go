package infrastructure

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"
	save "tradething/app/bn/infrastructure/future/save_side"
)

type ISavePositionSideBuilder interface {
	Get(ctx context.Context, position *position.Position) save.ISavePositionBySide
}

type ITradeBuilder interface {
	GetTradePosition(ctx context.Context, position_side string) position.IPosition
}
