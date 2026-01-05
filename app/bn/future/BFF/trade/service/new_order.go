package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/trade/domain"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *tradeService) NewOrder(ctx context.Context, order *domain.Order) error {
	history, err := s.positionHistoryService.Get(ctx, order.ClientId)
	if err != nil {
		return err
	}
	if history != nil {
		return errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	var isFromAdvancedPosition bool
	currentPosition, err := s.currentPositionService.Get(ctx, order.Symbol, order.AccountId, order.PositionSide)
	if err != nil {
		return err
	}
	if currentPosition != nil {
		return errors.New(appresponse.FoundCurrentPositionErrorCode)
	} else {
		advancedPosition, err := s.advancedPositionService.Get(ctx, order.ClientId)
		if err != nil {
			return err
		}
		if advancedPosition != nil {
			isFromAdvancedPosition = true
			order = order.NewOrderFrom(advancedPosition)
		}
	}

	subaccount, err := s.subaccountService.Get(ctx, order.AccountId)
	if err != nil {
		return err
	}
	if subaccount == nil {
		return errors.New(appresponse.SubAccountNotRegisteredErrorCode)
	}

	err = s.tradeAdaptor.NewOrder(ctx, order)
	if err != nil {
		return err
	}

	if order.IsBuyPosition() {
		err = s.currentPositionService.Upsert(ctx, order)
		if err != nil {
			return err
		}
		if isFromAdvancedPosition {
			err = s.advancedPositionService.Delete(ctx, order.ClientId)
			if err != nil {
				return err
			}
		}
	} else {
		err = s.currentPositionService.Delete(ctx, order.Symbol, order.AccountId, order.PositionSide)
		if err != nil {
			return err
		}
		err = s.positionHistoryService.Insert(ctx, order)
		if err != nil {
			return err
		}
	}

	return nil
}
