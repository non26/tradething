package service

import (
	"context"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) InvalidatePosition(
	ctx context.Context,
	request *svcfuture.InvalidatePositionServiceRequest,
) (*svchandlerres.InvalidatePositionHandlerResponse, error) {
	response := svchandlerres.InvalidatePositionHandlerResponse{
		Result: []svchandlerres.InvalidatePositionHandlerResponseData{},
	}
	for _, orderId := range request.OrderIds {
		dbHistory, err := b.bnFtHistoryTable.Get(ctx, orderId)
		if err != nil {
			return nil, err
		}
		if dbHistory.IsFound() {
			response.Result = append(response.Result, svchandlerres.InvalidatePositionHandlerResponseData{
				OrderId: orderId,
				Message: "position history found",
				Status:  "success",
			})
			continue
		}
		dbOpeningPosition, err := b.bnFtOpeningPositionTable.ScanWith(ctx, orderId)
		if err != nil {
			return nil, err
		}
		if dbOpeningPosition.IsFound() {
			bnreq := svcfuture.Position{}
			bnreq.SetClientOrderId(orderId)
			bnreq.SetPositionSide(dbOpeningPosition.PositionSide)
			bnreq.SetSide(dbOpeningPosition.Side)
			bnreq.SetEntryQuantity(dbOpeningPosition.AmountQ)
			bnreq.SetSymbol(dbOpeningPosition.Symbol)
			if b.positionSideType.IsLong(dbOpeningPosition.PositionSide) {
				bnreq.SetSide(b.sideType.Sell())
			} else {
				bnreq.SetSide(b.sideType.Buy())
			}

			_, err := b.PlaceSingleOrder(ctx, &bnreq)
			if err != nil {
				response.Result = append(response.Result, svchandlerres.InvalidatePositionHandlerResponseData{
					OrderId: orderId,
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
				ClientId:     orderId,
				Symbol:       dbOpeningPosition.Symbol,
				PositionSide: dbOpeningPosition.PositionSide,
			})
			if err != nil {
				// return nil, err
				continue
			}

			response.Result = append(response.Result, svchandlerres.InvalidatePositionHandlerResponseData{
				OrderId: orderId,
				Message: "success",
				Status:  "success",
			})
			continue
		} else {
			err := b.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
				ClientId: orderId,
			})
			if err != nil {
				response.Result = append(response.Result, svchandlerres.InvalidatePositionHandlerResponseData{
					OrderId: orderId,
					Message: err.Error(),
					Status:  "fail",
				})
				continue
			}
		}
	}
	return &response, nil
}
