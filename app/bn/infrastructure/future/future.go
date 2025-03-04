package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"
	closedomainsvc "tradething/app/bn/process/future/domain_service/close_position"
	tradedomainsvc "tradething/app/bn/process/future/domain_service/trade"
)

type ITrade interface {
	PlacePosition(ctx context.Context, position *position.Position) error
	// ClosePositionByClientId(ctx context.Context, clientId string) error
	// AccumulatePosition(ctx context.Context, position *Position) error
	// close position by client ids
	// close position by symbol
	// invalidate psoition
}

type ITradeLookUp interface {
	LookUp(ctx context.Context, position *position.Position) (*tradedomainsvc.TradeLookUp, error)
}

type ITradeSavePosition interface {
	Save(ctx context.Context, position *position.Position, lookup *tradedomainsvc.TradeLookUp) error
}

type IClosePositionLookup interface {
	ById(ctx context.Context, clientId string) (*closedomainsvc.ClsoePositionLookUp, error)
}
