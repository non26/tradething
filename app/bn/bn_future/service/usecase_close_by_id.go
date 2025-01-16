package service

import (
	"context"
	"errors"
	"log"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	handlerres "tradething/app/bn/bn_future/handler_response_model"
	model "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) CloseByClientIds(
	ctx context.Context,
	request *model.ClientIds,
) (*handlerres.CloseByClientIds, error) {
	closeOrders := handlerres.CloseByClientIds{
		Data: []handlerres.CloseByClientIdsData{},
	}
	for _, clientId := range request.GetCleintIds() {
		closeOrder := handlerres.CloseByClientIdsData{}
		positionHistory, err := b.bnFtHistoryTable.Get(ctx, clientId)
		if err != nil {
			return nil, errors.New("get history error " + err.Error())
		}
		if positionHistory.IsFound() {
			addCloseOrderData(&closeOrders, closeOrder, clientId, "fail", "no position history found")
			continue
		}
		openOrders, err := b.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			addCloseOrderData(&closeOrders, closeOrder, clientId, "fail", err.Error())
			continue
		}

		if !openOrders.IsFound() {
			addCloseOrderData(&closeOrders, closeOrder, clientId, "fail", "no open order found")
			continue
		}
		side := ""
		if b.positionSideType.IsLong(openOrders.PositionSide) {
			side = b.sideType.Sell()
		} else {
			side = b.sideType.Buy()
		}

		_, err = b.binanceService.PlaceSingleOrder(ctx, &bntradereq.PlacePosition{
			PositionSide:  openOrders.PositionSide,
			Side:          side,
			Symbol:        openOrders.Symbol,
			ClientOrderId: openOrders.ClientId,
			EntryQuantity: openOrders.AmountQ,
		})
		if err != nil {
			addCloseOrderData(&closeOrders, closeOrder, clientId, "fail", err.Error())
			continue
		}

		err = b.bnFtOpeningPositionTable.Delete(ctx, openOrders)
		if err != nil {
			log.Println("delete open order error", err)
		}
		err = b.bnFtHistoryTable.Insert(ctx, &dynamodbrepository.BnFtHistory{
			ClientId:     clientId,
			Symbol:       openOrders.Symbol,
			PositionSide: openOrders.PositionSide,
		})
		if err != nil {
			log.Println("insert history error", err)
		}
		addCloseOrderData(&closeOrders, closeOrder, clientId, "success", "close order success")
	}
	return &closeOrders, nil
}

func addCloseOrderData(clsoeOrders *handlerres.CloseByClientIds, closeOrder handlerres.CloseByClientIdsData, clientId string, status string, message string) {
	closeOrder.ClientId = clientId
	closeOrder.Status = status
	closeOrder.Message = message
	clsoeOrders.Data = append(clsoeOrders.Data, closeOrder)
}
