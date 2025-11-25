package service

import (
	"context"
	"tradething/app/bn/future/BFF/advanced_position/domain"
)

func (s *advancedPositionService) Get(ctx context.Context, clientId string) (*domain.AdvancedPosition, error) {
	positionExt, err := s.advancedService.GetAdvancedPosition().Get(ctx, clientId)
	if err != nil {
		return nil, err
	}

	advPosition := &domain.AdvancedPosition{
		ClientId:     positionExt.ClientId,
		Symbol:       positionExt.Symbol,
		PositionSide: positionExt.PositionSide,
		Side:         positionExt.Side,
		AmountB:      positionExt.AmountB,
		AccountId:    positionExt.AccountId,
	}

	return advPosition, nil
}
