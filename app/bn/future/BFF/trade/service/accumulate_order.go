package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/trade/domain"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *tradeService) AccumulateOrder(ctx context.Context, order *domain.Order) error {
	// check accum id
	history, err := s.positionHistoryService.Get(ctx, order.AccumID)
	if err != nil {
		return err
	}
	if history != nil {
		return errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	// check client id
	history, err = s.positionHistoryService.Get(ctx, order.ClientId)
	if err != nil {
		return err
	}
	if history != nil {
		return errors.New(appresponse.FoundPositionInHistoryErrorCode)
	}

	referencePosition, err := s.currentPositionService.ScanWithClientId(ctx, order.ClientId)
	if err != nil {
		return err
	}
	if referencePosition != nil {
		accumPosition, err := s.accumService.Get(ctx, order.ClientId)
		if err != nil {
			return err
		}
		if accumPosition.AccumID != order.AccumID {
			return errors.New(appresponse.NOTFOUNDACCUMLATIONCODE)
		}
		if accumPosition.ClientId != order.ClientId {
			return errors.New(appresponse.NOTFOUNDCLIENTIDCODE)
		}

		exceed, err := referencePosition.IsAmount1ExceedAmount2(referencePosition.AmountB, accumPosition.MaxAccum)
		if err != nil {
			return err
		}
		if exceed {
			err = s.accumService.Delete(ctx, order.ClientId)
			if err != nil {
				return err
			}
			accumPosition.ClientId = order.AccumID
			err = s.positionHistoryService.Insert(ctx, accumPosition)
			if err != nil {
				return err
			}
			return errors.New(appresponse.EXCEEDMAXACCUMULATIONCODE)
		}
		// refPositionAmountBeforeAccum := referencePosition.AmountB
		refPositionAmountAfterAccum, err := referencePosition.AddAmount(referencePosition.AmountB, order.Accum)
		if err != nil {
			return err
		}
		referencePosition.AmountB = order.Accum
		err = s.tradeAdaptor.NewOrder(ctx, referencePosition)
		if err != nil {
			return err
		}
		// newMaxAccum, err := order.AddAmount(order.MaxAccum, order.Accum)
		// if err != nil {
		// 	return err
		// }
		// update present accum
		accumPosition.AmountB = refPositionAmountAfterAccum
		err = s.accumService.Upsert(ctx, accumPosition)
		if err != nil {
			return err
		}
		referencePosition.AmountB = refPositionAmountAfterAccum
		err = s.currentPositionService.Upsert(ctx, referencePosition)
		if err != nil {
			return err
		}
	} else {
		return errors.New(appresponse.NotFoundOpeningPositionErrorCode)
	}

	return nil
}
