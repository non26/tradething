package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

func (b *binanceFutureService) InvalidatePosition(
	ctx context.Context,
	request *model.ClientIds,
) (*handlerres.InvalidatePosition, error) {
	response := handlerres.InvalidatePosition{
		Result: []handlerres.InvalidatePositionData{},
	}
	for _, clientId := range request.GetCleintIds() {
		dbHistory, err := b.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			return nil, errors.New("get history error " + err.Error())
		}
		if dbHistory.IsFound() {
			response.Result = append(response.Result, handlerres.InvalidatePositionData{
				OrderId: clientId,
				Message: "position history found",
				Status:  "success",
			})
			continue
		}

		dbOpeningPosition, err := b.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return nil, errors.New("get open order error " + err.Error())
		}
		if dbOpeningPosition.IsFound() {
			bnreq := model.Position{}
			bnreq.SetClientId(clientId)
			bnreq.SetPositionSide(dbOpeningPosition.PositionSide)
			bnreq.SetSide(dbOpeningPosition.Side)
			bnreq.SetEntryQuantity(dbOpeningPosition.AmountB)
			bnreq.SetSymbol(dbOpeningPosition.Symbol)
			if utils.IsLongPosition(dbOpeningPosition.PositionSide) {
				bnreq.SetSide(bnconstant.SELL)
			} else {
				bnreq.SetSide(bnconstant.BUY)
			}

			_, err := b.PlaceSingleOrder(ctx, &bnreq)
			if err != nil {
				response.Result = append(response.Result, handlerres.InvalidatePositionData{
					OrderId: clientId,
					Message: err.Error(),
					Status:  "fail",
				})
				continue
			}

			err = b.bnFtOpeningPositionTable.Delete(ctx, dbOpeningPosition)
			if err != nil {
				// return nil, err
				continue
			}

			err = b.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
				ClientId:     clientId,
				Symbol:       dbOpeningPosition.Symbol,
				PositionSide: dbOpeningPosition.PositionSide,
			})
			if err != nil {
				// return nil, err
				continue
			}

			response.Result = append(response.Result, handlerres.InvalidatePositionData{
				OrderId: clientId,
				Message: "success",
				Status:  "success",
			})
			continue
		}

		dbAdvancePosition, err := b.bnFtAdvancedPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			continue
		}
		if dbAdvancePosition.IsFound() {
			err := b.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
				ClientId: clientId,
			})
			if err != nil {
				continue
			}
			response.Result = append(response.Result, handlerres.InvalidatePositionData{
				OrderId: clientId,
				Message: "success",
				Status:  "success",
			})
			continue
		}

		err = b.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
			ClientId: clientId,
		})
		if err != nil {
			response.Result = append(response.Result, handlerres.InvalidatePositionData{
				OrderId: clientId,
				Message: err.Error(),
				Status:  "fail",
			})
			continue
		}

	}
	return &response, nil
}
