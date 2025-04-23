package infrastructure

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"

	future "tradething/app/bn/infrastructure/future"
	builder "tradething/app/bn/infrastructure/future/builder"
	domainAdvPositionSvc "tradething/app/bn/process/future/domain_service/advanced_position"
	domainCryptoSvc "tradething/app/bn/process/future/domain_service/crypto"
	domainTradeSvc "tradething/app/bn/process/future/domain_service/trade"
)

type savePosition struct {
	queryPosition builder.ISavePositionBuilder
}

func NewSavePosition(queryPosition builder.ISavePositionBuilder) future.ITradeSavePosition {
	return &savePosition{queryPosition}
}

func (s *savePosition) Save(ctx context.Context, position *position.Position, tradeLookup *domainTradeSvc.TradeLookUp, cryptoLookup *domainCryptoSvc.CryptoLookUp, advPositionLookUp *domainAdvPositionSvc.AdvancedPositionLookUp) error {
	savePositionBySide := s.queryPosition.Get(ctx, position)
	return savePositionBySide.Save(ctx, position, tradeLookup, cryptoLookup, advPositionLookUp)
}
