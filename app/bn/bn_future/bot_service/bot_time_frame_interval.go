package bnfuture

import (
	"context"
	"fmt"
	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"
)

func (b *botService) TimeIntervalSemiBotService(
	ctx context.Context,
	req *bothandlerreq.TradeTimeIntervalBinanceFutureRequest,
) (*bothandlerres.TradeTimeIntervalBinanceFutureResponse, error) {
	openingOrder, err := b.bn_repository.GetOpenOrderBySymbol(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}
	currentSymbolQoute, err := b.bn_repository.GetQouteUSDT(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	prev_side := openingOrder.Side
	prev_position_side := openingOrder.PositionSide

	var openNewOrderReq *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	var closeOlderOrderReq *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	if !openingOrder.IsEmpty() {
		if b.side.IsBuy(prev_side) && b.position_side.IsLong(prev_position_side) {
			if b.position_side.IsLong(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Buy())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
			} else if b.position_side.IsShort(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Sell())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
			}
			closeOlderOrderReq = closeOlderOrderReq.New()
			closeOlderOrderReq.SetPositionSide(b.position_side.Long())
			closeOlderOrderReq.SetSide(b.side.Sell())
			closeOlderOrderReq.SetType(b.order_type.Market())
			closeOlderOrderReq.SetEntryQuantity(openingOrder.AmountQ)
			closeOlderOrderReq.SetSymbol(req.Symbol)
			closeOlderOrderReq.SetClientOrderId(openNewOrderReq.ClientOrderId)
		} else if b.side.IsSell(prev_side) && b.position_side.IsShort(prev_position_side) {
			if b.position_side.IsLong(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Buy())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
			} else if b.position_side.IsShort(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Sell())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
			}
			closeOlderOrderReq = closeOlderOrderReq.New()
			closeOlderOrderReq.SetPositionSide(b.position_side.Short())
			closeOlderOrderReq.SetSide(b.side.Buy())
			closeOlderOrderReq.SetType(b.order_type.Market())
			closeOlderOrderReq.SetEntryQuantity(openingOrder.AmountQ)
			closeOlderOrderReq.SetSymbol(openingOrder.Symbol)
			closeOlderOrderReq.SetClientOrderId(openNewOrderReq.ClientOrderId)
		}
	} else {
		if b.position_side.IsLong(req.PositionSide) {
			openNewOrderReq = openNewOrderReq.New()
			openNewOrderReq.SetPositionSide(req.PositionSide)
			openNewOrderReq.SetSide(b.side.Buy())
			openNewOrderReq.SetType(b.order_type.Market())
			openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
			openNewOrderReq.SetSymbol(req.Symbol)
			openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
		} else {
			openNewOrderReq = openNewOrderReq.New()
			openNewOrderReq.SetPositionSide(req.PositionSide)
			openNewOrderReq.SetSide(b.side.Sell())
			openNewOrderReq.SetType(b.order_type.Market())
			openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
			openNewOrderReq.SetSymbol(req.Symbol)
			openNewOrderReq.SetDefaultClientOrderId(currentSymbolQoute.GetNextCounting().String())
		}
	}

	if closeOlderOrderReq != nil {
		_, err := b.bn_service.PlaceSingleOrder(ctx, closeOlderOrderReq)
		if err != nil {
			return nil, err
		}
	}
	if openNewOrderReq != nil {
		if req.LeverageLevel != "" {
			_, err := b.bn_service.SetNewLeverage(ctx, &bnserivcemodelreq.SetLeverageBinanceServiceRequest{
				Symbol:   openNewOrderReq.Symbol,
				Leverage: req.LeverageLevel,
			})
			if err != nil {
				return nil, err
			}
		}

		_, err = b.bn_service.PlaceSingleOrder(ctx, openNewOrderReq)
		if err != nil {
			return nil, err
		}
		err = b.bn_repository.UpdateCountingSymbolQouteUSDT(ctx, currentSymbolQoute)
		if err != nil {
			return nil, err
		}
		if !openingOrder.IsEmpty() {
			err = b.bn_repository.DeleteOpenOrderBySymbol(ctx, req.Symbol)
			if err != nil {
				return nil, err
			}
		}
		err = b.bn_repository.NewOpenOrder(
			ctx,
			req.ToBnFutureOpeningPositionEntity(
				openNewOrderReq.Side,
				req.LeverageLevel,
				currentSymbolQoute.GetNextCounting().String()))
		if err != nil {
			return nil, err
		}
		return &bothandlerres.TradeTimeIntervalBinanceFutureResponse{
			Message: "Suucess",
		}, nil
	}

	return &bothandlerres.TradeTimeIntervalBinanceFutureResponse{
		Message: "Suucess But No Order Made",
	}, nil
}
