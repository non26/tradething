package infrastructure

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"
	"tradething/app/bn/infrastructure/future/save"
	domainservice "tradething/app/bn/process/future/domain_service/trade"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

type ISavePositionBuilder interface {
	Get(ctx context.Context, position *position.Position) save.ISavePositionBySide
}

type savePositionBuilder struct {
	saveBuyPosition  save.ISavePositionBySide
	saveSellPosition save.ISavePositionBySide
}

func NewSavePositionBuilder(saveBuyPosition save.ISavePositionBySide, saveSellPosition save.ISavePositionBySide) ISavePositionBuilder {
	return &savePositionBuilder{saveBuyPosition, saveSellPosition}
}

func (s *savePositionBuilder) Get(ctx context.Context, position *position.Position) save.ISavePositionBySide {
	if utils.IsBuyPosition(position.Side, position.PositionSide) {
		return s.saveBuyPosition
	}
	return s.saveSellPosition
}

type savePosition struct {
	queryPosition ISavePositionBuilder
}

func NewSavePosition(queryPosition ISavePositionBuilder) ISavePosition {
	return &savePosition{queryPosition}
}

func (s *savePosition) Save(ctx context.Context, position *position.Position, lookup *domainservice.LookUp) error {
	savePositionBySide := s.queryPosition.Get(ctx, position)
	return savePositionBySide.Save(ctx, position, lookup)
}
