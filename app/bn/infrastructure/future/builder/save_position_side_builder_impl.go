package infrastructure

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"
	save "tradething/app/bn/infrastructure/future/save_side"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

type savePositionSideBuilder struct {
	saveBuySide  save.ISavePositionBySide
	saveSellSide save.ISavePositionBySide
}

func NewSavePositionSideBuilder(saveBuyPosition save.ISavePositionBySide, saveSellPosition save.ISavePositionBySide) ISavePositionSideBuilder {
	return &savePositionSideBuilder{saveBuyPosition, saveSellPosition}
}

func (s *savePositionSideBuilder) Get(ctx context.Context, position *position.Position) save.ISavePositionBySide {
	if utils.IsBuyPosition(position.Side, position.PositionSide) {
		return s.saveBuySide
	}
	return s.saveSellSide
}
