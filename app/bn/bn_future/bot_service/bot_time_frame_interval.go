package bnfuture

// import (
// 	"context"
// 	"fmt"
// 	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
// 	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
// 	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"

// 	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
// )

// // TODO local testing for interacting with bn_future_qoute_usdt table
// // TODO scenario 1 if the symbol is not exist in the table, insert new symbol and insert new counting
// // TODO scenario 2 if the symbol is exist in the table, update the counting
// func (b *botService) TimeIntervalSemiBotService(
// 	ctx context.Context,
// 	req *bothandlerreq.TradeTimeIntervalBinanceFutureRequest,
// ) (*bothandlerres.TradeTimeIntervalBinanceFutureResponse, error) {
// 	openingOrder, err := b.bn_repository.GetOpenOrderBySymbol(ctx, req.Symbol)
// 	if err != nil {
// 		return nil, err
// 	}

// 	currentSymbolQoute, err := b.bn_repository.GetQouteUSDT(ctx, req.Symbol)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !currentSymbolQoute.IsExist() {
// 		currentSymbolQoute = &dynamodbmodel.BinanceFutureQouteUSDT{}
// 		currentSymbolQoute.SetCounting(-1)
// 		currentSymbolQoute.SetSymbol(req.Symbol)
// 	}

// 	var openNewOrderReq *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
// 	var closeOlderOrderReq *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
// 	if !openingOrder.IsEmpty() {
// 		if b.side.IsBuy(openingOrder.Side) && b.position_side.IsLong(openingOrder.PositionSide) {
// 			if b.position_side.IsLong(req.PositionSide) {
// 				openNewOrderReq = openNewOrderReq.New()
// 				openNewOrderReq.SetPositionSide(req.PositionSide)
// 				openNewOrderReq.SetSide(b.side.Buy())
// 				openNewOrderReq.SetType(b.order_type.Market())
// 				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
// 				openNewOrderReq.SetSymbol(req.Symbol)
// 				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
// 			} else if b.position_side.IsShort(req.PositionSide) {
// 				openNewOrderReq = openNewOrderReq.New()
// 				openNewOrderReq.SetPositionSide(req.PositionSide)
// 				openNewOrderReq.SetSide(b.side.Sell())
// 				openNewOrderReq.SetType(b.order_type.Market())
// 				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
// 				openNewOrderReq.SetSymbol(req.Symbol)
// 				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
// 			}
// 			closeOlderOrderReq = closeOlderOrderReq.New()
// 			closeOlderOrderReq.SetPositionSide(b.position_side.Long())
// 			closeOlderOrderReq.SetSide(b.side.Sell())
// 			closeOlderOrderReq.SetType(b.order_type.Market())
// 			closeOlderOrderReq.SetEntryQuantity(openingOrder.AmountQ)
// 			closeOlderOrderReq.SetSymbol(req.Symbol)
// 			closeOlderOrderReq.SetClientOrderId(openNewOrderReq.ClientOrderId)
// 		} else if b.side.IsSell(openingOrder.Side) && b.position_side.IsShort(openingOrder.PositionSide) {
// 			if b.position_side.IsLong(req.PositionSide) {
// 				openNewOrderReq = openNewOrderReq.New()
// 				openNewOrderReq.SetPositionSide(req.PositionSide)
// 				openNewOrderReq.SetSide(b.side.Buy())
// 				openNewOrderReq.SetType(b.order_type.Market())
// 				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
// 				openNewOrderReq.SetSymbol(req.Symbol)
// 				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
// 			} else if b.position_side.IsShort(req.PositionSide) {
// 				openNewOrderReq = openNewOrderReq.New()
// 				openNewOrderReq.SetPositionSide(req.PositionSide)
// 				openNewOrderReq.SetSide(b.side.Sell())
// 				openNewOrderReq.SetType(b.order_type.Market())
// 				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
// 				openNewOrderReq.SetSymbol(req.Symbol)
// 				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
// 			}
// 			closeOlderOrderReq = closeOlderOrderReq.New()
// 			closeOlderOrderReq.SetPositionSide(b.position_side.Short())
// 			closeOlderOrderReq.SetSide(b.side.Buy())
// 			closeOlderOrderReq.SetType(b.order_type.Market())
// 			closeOlderOrderReq.SetEntryQuantity(openingOrder.AmountQ)
// 			closeOlderOrderReq.SetSymbol(openingOrder.Symbol)
// 			closeOlderOrderReq.SetClientOrderId(openNewOrderReq.ClientOrderId)
// 		}
// 	} else {
// 		if b.position_side.IsLong(req.PositionSide) {
// 			openNewOrderReq = openNewOrderReq.New()
// 			openNewOrderReq.SetPositionSide(req.PositionSide)
// 			openNewOrderReq.SetSide(b.side.Buy())
// 			openNewOrderReq.SetType(b.order_type.Market())
// 			openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
// 			openNewOrderReq.SetSymbol(req.Symbol)
// 			openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
// 		} else {
// 			openNewOrderReq = openNewOrderReq.New()
// 			openNewOrderReq.SetPositionSide(req.PositionSide)
// 			openNewOrderReq.SetSide(b.side.Sell())
// 			openNewOrderReq.SetType(b.order_type.Market())
// 			openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
// 			openNewOrderReq.SetSymbol(req.Symbol)
// 			openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
// 		}
// 	}

// 	if closeOlderOrderReq != nil {
// 		_, err := b.bn_service.PlaceSingleOrder(ctx, closeOlderOrderReq)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if openNewOrderReq != nil {
// 		// if req.LeverageLevel != "" {
// 		// 	_, err := b.bn_service.SetNewLeverage(ctx, &bnserivcemodelreq.SetLeverageBinanceServiceRequest{
// 		// 		Symbol:   openNewOrderReq.Symbol,
// 		// 		Leverage: req.LeverageLevel,
// 		// 	})
// 		// 	if err != nil {
// 		// 		return nil, err
// 		// 	}
// 		// }

// 		// TODO check the transaction when interacting with the database
// 		if currentSymbolQoute.GetCounting() < 0 {
// 			err = b.bn_repository.InsertNewSymbolUSDT(ctx, req.Symbol)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			err = b.bn_repository.UpdateCountingSymbolQouteUSDT(
// 				ctx, req.ToBnFutureQouteUSDTEntity(currentSymbolQoute.GetCounting()))
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		if !openingOrder.IsEmpty() {
// 			err = b.bn_repository.DeleteOpenOrderBySymbol(ctx, req.Symbol)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		err = b.bn_repository.NewOpenOrder(
// 			ctx,
// 			req.ToBnFutureOpeningPositionEntity(
// 				openNewOrderReq.Side,
// 				req.LeverageLevel,
// 				openNewOrderReq.ClientOrderId))
// 		if err != nil {
// 			return nil, err
// 		}

// 		// TODO change the leverage level before making the order
// 		_, err = b.bn_service.PlaceSingleOrder(ctx, openNewOrderReq)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return &bothandlerres.TradeTimeIntervalBinanceFutureResponse{
// 			Message: "Suucess",
// 		}, nil
// 	}

// 	return &bothandlerres.TradeTimeIntervalBinanceFutureResponse{
// 		Message: "Suucess But No Order Made",
// 	}, nil
// }
