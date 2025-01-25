package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response"
	svcmodels "tradething/app/bn/bn_future/service_model"

	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
)

func (b *binanceFutureService) ValidateAdavancedPosition(ctx context.Context, clientIds *svcmodels.ClientIds) (*handlerres.ValidatePosition, serviceerror.IError) {
	response := handlerres.ValidatePosition{
		Data: []handlerres.ValidatePositionData{},
	}
	for _, clientId := range clientIds.GetCleintIds() {
		dbHistory, err := b.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			addValidatePositionData(&response, clientId, "fail", err.Error())
			continue
		}
		if dbHistory.IsFound() {
			addValidatePositionData(&response, clientId, "success", "history found")
			continue
		}

		dbOpeningPosition, err := b.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			addValidatePositionData(&response, clientId, "fail", err.Error())
			continue
		}
		if dbOpeningPosition.IsFound() {
			addValidatePositionData(&response, clientId, "success", "opening position found")
			continue
		}

		dbAdvancedPosition, err := b.bnFtAdvancedPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			addValidatePositionData(&response, clientId, "fail", err.Error())
			continue
		}
		if dbAdvancedPosition.IsFound() {
			request := svcmodels.Position{}
			request.SetSymbol(dbAdvancedPosition.GetSymbol())
			request.SetPositionSide(dbAdvancedPosition.GetPositionSide())
			request.SetEntryQuantity(dbAdvancedPosition.GetAmountB())
			request.SetSide(dbAdvancedPosition.GetSide())
			_, err := b.openPosition(ctx, &request)
			if err != nil {
				addValidatePositionData(&response, clientId, "fail", err.Error())
				continue
			}
			addValidatePositionData(&response, clientId, "success", "opening position found")
		} else {
			addValidatePositionData(&response, clientId, "success", "advanced position not found")
		}
	}
	return &response, nil
}

func addValidatePositionData(response *handlerres.ValidatePosition, clientId string, status string, message string) {
	response.Data = append(response.Data, handlerres.ValidatePositionData{
		ClientId: clientId,
		Status:   status,
		Message:  message,
	})
}
