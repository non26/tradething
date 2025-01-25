package service

import (
	"context"
	"errors"
	svcmodels "tradething/app/bn/bn_future/service_model"
)

func (b *binanceFutureService) ValidateAdavancedPosition(ctx context.Context, clientIds svcmodels.ClientIds) error {

	for _, clientId := range clientIds.GetCleintIds() {
		dbHistory, err := b.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			return errors.New("get history table error")
		}
		if dbHistory.IsFound() {
			continue
		}

		dbOpeningPosition, err := b.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return errors.New("scan openign position table error")
		}
		if dbOpeningPosition.IsFound() {
			continue
		}

		dbAdvancedPosition, err := b.bnFtAdvancedPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return errors.New("scan advanced position table error")
		}
		if dbAdvancedPosition.IsFound() {
			request := svcmodels.Position{}
			request.SetSymbol(dbAdvancedPosition.GetSymbol())
			request.SetPositionSide(dbAdvancedPosition.GetPositionSide())
			request.SetEntryQuantity(dbAdvancedPosition.GetAmountB())
			request.SetSide(dbAdvancedPosition.GetSide())
			_, err := b.openPosition(ctx, &request)
			if err != nil {
				return errors.New("open position error")
			}
		}
	}
	return nil
}
