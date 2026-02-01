package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/trade/domain"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
	"github.com/shopspring/decimal"
)

func (s *tradeService) NewOrder(ctx context.Context, order *domain.Order) error {
	history, err := s.positionHistoryService.Get(ctx, order.ClientId)
	if err != nil {
		return err
	}
	if history != nil {
		return errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	isFromAdvancedPosition := order.IsAdvancedPosition()
	var currentPosition *domain.Order
	if !isFromAdvancedPosition {
		currentPosition, err = s.currentPositionService.Get(ctx, order.Symbol, order.AccountId, order.PositionSide)
		if err != nil {
			return err
		}
	}
	if currentPosition != nil {
		if currentPosition.ClientId == order.ClientId {
			return errors.New(appresponse.FoundCurrentPositionErrorCode)
		}
		_currentPositionAmountB, err := decimal.NewFromString(currentPosition.AmountB)
		if err != nil {
			return err
		}
		_reqAmountB, err := decimal.NewFromString(order.AmountB)
		if err != nil {
			return err
		}
		if !_reqAmountB.Equal(_currentPositionAmountB) {
			return errors.New(appresponse.NotFoundOpeningPositionErrorCode)
		}
		if currentPosition.PositionSide != order.PositionSide {
			return errors.New(appresponse.FoundCurrentPositionErrorCode)
		}
		if currentPosition.Symbol != order.Symbol {
			return errors.New(appresponse.NotFoundOpeningPositionErrorCode)
		}
	} else {
		advancedPosition, err := s.advancedPositionService.Get(ctx, order.ClientId)
		if err != nil {
			return err
		}
		if advancedPosition != nil {
			order = order.NewOrderFrom(advancedPosition)
		} else {
			isFromAdvancedPosition = false
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
