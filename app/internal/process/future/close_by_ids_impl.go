package process

import (
	"context"
	response "tradething/app/internal/handlers/future/res"
	"tradething/app/internal/process/future/domain"
)

func (f *future) ClosePositionByClientIds(ctx context.Context, clientIds []string) (response *response.CloseByClientIds, err error) {
	for _, clientId := range clientIds {
		bnHistory, err := f.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			response.AddWithData("error", "error", "error", "error", clientId)
			continue
		}
		if bnHistory.IsFound() {
			response.AddWithData("error", "error", bnHistory.Symbol, bnHistory.PositionSide, clientId)
			continue
		}

		bnOpening, err := f.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return nil, err
		}
		if !bnOpening.IsFound() {
			response.AddWithData("error", "error", "error", "error", clientId)
			continue
		}

		position := domain.Position{}
		position.SetClientId(clientId)
		position.SetSymbol(bnOpening.Symbol)
		position.SetPositionSide(bnOpening.PositionSide)
		position.SetSide(bnOpening.Side)
		position.SetEntryQuantity(bnOpening.AmountB)

		err = f.infraFuture.PlacePosition(ctx, position.ToInfraPosition())
		if err != nil {
			response.AddWithData("error", "error", bnOpening.Symbol, bnHistory.PositionSide, clientId)
			continue
		}
		response.AddWithData("success", "success", bnOpening.Symbol, bnOpening.PositionSide, clientId)
	}

	return response, nil
}
