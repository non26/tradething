package bnfuture

// import (
// 	"context"
// 	"errors"
// 	handlerres "tradething/app/bn/bn_future/bot_handler_response_model"
// 	bnbotsvcreq "tradething/app/bn/bn_future/bot_service_model"

// 	bntrademodel "tradething/app/bn/bn_future/bnservice_request_model/trade"

// 	bndynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
// )

// func (b *botService) InvalidateBot(ctx context.Context, req *bnbotsvcreq.InvalidateBot) (*handlerres.InvalidateBotHandlerResponse, error) {

// 	history, err := b.bnFtHistoryTable.Get(ctx, req.BotOrderId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if history.IsFound() {
// 		return nil, errors.New("bot already closed")
// 	}

// 	currentPosition, err := b.repository.GetOpenOrderBySymbolAndPositionSide(ctx, &bndynamodbmodel.BnFtOpeningPosition{
// 		Symbol:       req.Symbol,
// 		PositionSide: req.PositionSide,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var side string
// 	if currentPosition.PositionSide == b.positionSideType.Long() {
// 		side = b.sideType.Sell()
// 	} else {
// 		side = b.sideType.Buy()
// 	}

// 	if currentPosition.IsFound() {
// 		_, err = b.binanceService.PlaceSingleOrder(ctx, &bntrademodel.PlaceSignleOrderBinanceServiceRequest{
// 			Symbol:        currentPosition.Symbol,
// 			PositionSide:  currentPosition.PositionSide,
// 			Side:          side,
// 			Type:          b.orderType.Market(),
// 			EntryQuantity: currentPosition.AmountQ,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = b.repository.DeleteBotOnRun(ctx, &bndynamodbmodel.BnFtBotOnRun{
// 			BotID:      req.BotId,
// 			BotOrderID: req.BotOrderId,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = b.repository.InsertHistory(ctx, &bndynamodbmodel.BnFtHistory{
// 			ClientId:     req.BotOrderId,
// 			Symbol:       currentPosition.Symbol,
// 			PositionSide: currentPosition.PositionSide,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return nil, nil
// }
