package service

import (
	"context"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

func (b *binanceFutureService) InvalidatePosition(
	ctx context.Context,
	request *model.ClientIds,
) (*handlerres.InvalidatePosition, serviceerror.IError) {
	response := handlerres.InvalidatePosition{
		Result: []handlerres.InvalidatePositionData{},
	}
	for _, clientId := range request.GetCleintIds() {
		dbHistory, err := b.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, err)
		}
		if dbHistory.IsFound() {
			addInValidatePositionData(&response, clientId, "position history found", "success")
			continue
		}

		dbOpeningPosition, err := b.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
		}
		if dbOpeningPosition.IsFound() {
			bnreq := model.Position{}
			bnreq.SetClientId(clientId)
			bnreq.SetPositionSide(dbOpeningPosition.PositionSide)
			bnreq.SetSide(dbOpeningPosition.Side)
			bnreq.SetEntryQuantity(dbOpeningPosition.AmountB)
			bnreq.SetSymbol(dbOpeningPosition.Symbol)
			bnreq.SetSide(utils.ToSellSideBy(dbOpeningPosition.PositionSide))

			_, svcerr := b.PlaceSingleOrder(ctx, &bnreq)
			if svcerr != nil {
				addInValidatePositionData(&response, clientId, svcerr.Error(), "fail")
				continue
			}

			err = b.bnFtOpeningPositionTable.Delete(ctx, dbOpeningPosition)
			if err != nil {
				continue
			}

			err = b.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
				ClientId:     clientId,
				Symbol:       dbOpeningPosition.Symbol,
				PositionSide: dbOpeningPosition.PositionSide,
			})
			if err != nil {
				continue
			}
			addInValidatePositionData(&response, clientId, "success", "success")
		}
	}
	return &response, nil
}

func addInValidatePositionData(response *handlerres.InvalidatePosition, clientId string, message string, status string) {
	response.Result = append(response.Result, handlerres.InvalidatePositionData{
		OrderId: clientId,
		Message: message,
		Status:  status,
	})
}
