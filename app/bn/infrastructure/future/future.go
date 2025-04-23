package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"
	domainAdvPositionSvc "tradething/app/bn/process/future/domain_service/advanced_position"
	domainCryptoSvc "tradething/app/bn/process/future/domain_service/crypto"
	domainTradeSvc "tradething/app/bn/process/future/domain_service/trade"
	// tradedomainsvc "tradething/app/bn/process/future/domain_service/trade"
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
	LookUp(ctx context.Context, position *position.Position) (*domainTradeSvc.TradeLookUp, error)
}

type ITradeSavePosition interface {
	Save(ctx context.Context, position *position.Position, tradeLookup *domainTradeSvc.TradeLookUp, cryptoLookup *domainCryptoSvc.CryptoLookUp, advPositionLookUp *domainAdvPositionSvc.AdvancedPositionLookUp) error
}

type IAdvancedPositionLookup interface {
	LookUpByClientId(ctx context.Context, clientId string) (*domainAdvPositionSvc.AdvancedPositionLookUp, error)
}

type ICryptoLookUp interface {
	LookUpBySymbol(ctx context.Context, symbol string, positionSide string) (*domainCryptoSvc.CryptoLookUp, error)
}
