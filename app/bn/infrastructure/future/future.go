package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"
	domainservice "tradething/app/bn/process/future/domain_service"
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
	LookUp(ctx context.Context, position *position.Position) (*domainservice.LookUp, error)
}

type ISavePosition interface {
	Save(ctx context.Context, position *position.Position, lookup *domainservice.LookUp) error
}
