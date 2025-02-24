package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"
)

type ITrade interface {
	PlacePosition(ctx context.Context, position *position.Position) error
	// ClosePositionByClientId(ctx context.Context, clientId string) error
	// AccumulatePosition(ctx context.Context, position *Position) error
	// close position by client ids
	// close position by symbol
	// invalidate psoition
}

type ITradePosition interface {
	GetPosition(ctx context.Context, position_side string) position.IPosition
}
