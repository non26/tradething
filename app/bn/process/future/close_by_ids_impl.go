package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	"tradething/app/bn/process/future/domain"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func (f *future) ClosePositionByClientIds(ctx context.Context, clientIds []string) (responses *response.CloseByClientIds, err error) {
	responses = &response.CloseByClientIds{}
	for _, clientId := range clientIds {
		bnHistory, err := f.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			responses.AddWithData("error", "error", "error", "error", clientId)
			continue
		}
		if bnHistory.IsFound() {
			responses.AddWithData("error", "error", bnHistory.Symbol, bnHistory.PositionSide, clientId)
			continue
		}

		bnOpening, err := f.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return nil, err
		}
		if !bnOpening.IsFound() {
			responses.AddWithData("error", "error", "error", "error", clientId)
			continue
		}

		position := domain.Position{}
		position.SetClientId(clientId)
		position.SetSymbol(bnOpening.Symbol)
		position.SetPositionSide(bnOpening.PositionSide)
		if bnOpening.PositionSide == bnconstant.LONG {
			position.SetSide(bnconstant.SELL)
		} else {
			position.SetSide(bnconstant.BUY)
		}
		position.SetEntryQuantity(bnOpening.AmountB)

		err = f.infraFuture.PlacePosition(ctx, position.ToInfraPosition())
		if err != nil {
			responses.AddWithData("error", "error", bnOpening.Symbol, bnHistory.PositionSide, clientId)
			continue
		}
		responses.AddWithData("success", "success", bnOpening.Symbol, bnOpening.PositionSide, clientId)
	}

	return responses, nil
}
