package service

import (
	"context"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	svchandlerres "tradething/app/bn/bn_future/handler_response_model"
	svcfuture "tradething/app/bn/bn_future/service_model"

	dynamodbrepository "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) CloseByClientIds(
	ctx context.Context,
	request *svcfuture.CloseByClientIdServiceRequest,
) (*svchandlerres.CloseByClientIdsHandlerResponse, error) {
	closeOrders := svchandlerres.CloseByClientIdsHandlerResponse{
		Data: []svchandlerres.CloseByClientIdsHandlerResponseData{},
	}
	for _, clientId := range request.GetClientIds() {
		closeOrder := svchandlerres.CloseByClientIdsHandlerResponseData{}
		positionHistory, err := b.repository.GetHistoryByClientID(ctx, clientId)
		if err != nil {
			return nil, err
		}
		if positionHistory.IsFound() {
			addCloseOrderData(&closeOrders, closeOrder, clientId, "fail", "no position history found")
			continue
		}
		openOrders, err := b.repository.GetOpenOrderByClientID(ctx, clientId)
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

		_, err = b.binanceService.PlaceSingleOrder(ctx, &bntradereq.PlaceSignleOrderBinanceServiceRequest{
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

		b.repository.DeleteOpenOrderByKey(ctx, openOrders.GetKeyByPositionSideAndSymbol())
		b.repository.InsertHistory(ctx, &dynamodbrepository.BinanceFutureHistory{
			ClientId:     clientId,
			Symbol:       openOrders.Symbol,
			PositionSide: openOrders.PositionSide,
		})
		addCloseOrderData(&closeOrders, closeOrder, clientId, "success", "close order success")
	}
	return &closeOrders, nil
}

func addCloseOrderData(clsoeOrders *svchandlerres.CloseByClientIdsHandlerResponse, closeOrder svchandlerres.CloseByClientIdsHandlerResponseData, clientId string, status string, message string) {
	closeOrder.ClientId = clientId
	closeOrder.Status = status
	closeOrder.Message = message
	clsoeOrders.Data = append(clsoeOrders.Data, closeOrder)
}
