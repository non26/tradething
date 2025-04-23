package infrastructure

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"
	save "tradething/app/bn/infrastructure/future/save"

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
