package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"
	domainservice "tradething/app/bn/process/future/domain_service"
	advancedpositiondomainsvc "tradething/app/bn/process/future/domain_service/advanced_position"
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

// type IClosePositionLookup interface {
// 	ById(ctx context.Context, clientId string) (*closedomainsvc.ClsoePositionLookUp, error)
// }

type IAdvancedPositionLookup interface {
	LookUpByClientId(ctx context.Context, clientId string) (*advancedpositiondomainsvc.AdvancedPositionLookUp, error)
}

type ICryptoLookUp interface {
	LookUpBySymbol(ctx context.Context, symbol string, positionSide string) (*domainservice.LookUpSymbol, error)
}
