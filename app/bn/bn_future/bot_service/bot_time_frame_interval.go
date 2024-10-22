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
				openNewOrderReq.SetClientOrderId(req.CurrentClientId)
			} else if b.position_side.IsShort(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Sell())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetClientOrderId(req.CurrentClientId)
			}
			closeOlderOrderReq = closeOlderOrderReq.New()
			closeOlderOrderReq.SetPositionSide(b.position_side.Long())
			closeOlderOrderReq.SetSide(b.side.Sell())
			closeOlderOrderReq.SetType(b.order_type.Market())
			closeOlderOrderReq.SetEntryQuantity(openingOrder.AmountQ)
			closeOlderOrderReq.SetSymbol(req.Symbol)
			closeOlderOrderReq.SetClientOrderId(req.PrevClientId)
		} else if b.side.IsSell(prev_side) && b.position_side.IsShort(prev_position_side) {
			if b.position_side.IsLong(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Buy())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetClientOrderId(req.CurrentClientId)
			} else if b.position_side.IsShort(req.PositionSide) {
				openNewOrderReq = openNewOrderReq.New()
				openNewOrderReq.SetPositionSide(req.PositionSide)
				openNewOrderReq.SetSide(b.side.Sell())
				openNewOrderReq.SetType(b.order_type.Market())
				openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				openNewOrderReq.SetSymbol(req.Symbol)
				openNewOrderReq.SetClientOrderId(req.CurrentClientId)
			}
			closeOlderOrderReq = closeOlderOrderReq.New()
			closeOlderOrderReq.SetPositionSide(b.position_side.Short())
			closeOlderOrderReq.SetSide(b.side.Buy())
			closeOlderOrderReq.SetType(b.order_type.Market())
			closeOlderOrderReq.SetEntryQuantity(openingOrder.AmountQ)
			closeOlderOrderReq.SetSymbol(openingOrder.Symbol)
			closeOlderOrderReq.SetClientOrderId(req.PrevClientId)
		}
	} else {
		openNewOrderReq = openNewOrderReq.New()
		openNewOrderReq.SetPositionSide(req.PositionSide)
		openNewOrderReq.SetSide(b.side.Buy())
		openNewOrderReq.SetType(b.order_type.Market())
		openNewOrderReq.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
		openNewOrderReq.SetSymbol(req.Symbol)
		openNewOrderReq.SetClientOrderId(req.CurrentClientId)
	}

	if closeOlderOrderReq != nil {
		_, err := b.bn_service.PlaceSingleOrder(ctx, closeOlderOrderReq)
		if err != nil {
			return nil, err
		}
	}
	if openNewOrderReq != nil {
		// if req.LeverageLevel != "" {
		// 	_, err := b.bn_service.SetNewLeverage(ctx, &bnserivcemodelreq.SetLeverageBinanceServiceRequest{
		// 		Symbol:   openNewOrderReq.Symbol,
		// 		Leverage: req.LeverageLevel,
		// 	})
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// }

		_, err = b.bn_service.PlaceSingleOrder(ctx, openNewOrderReq)
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
