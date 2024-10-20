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
	query_request := req.ToQueryOrderBinanceServiceRequest()
	q_data, err := b.bn_service.QueryOrder(
		ctx,
		query_request)
	if err != nil {
		return nil, err
	}

	prev_side := q_data.Side
	prev_position_side := q_data.PositionSide
	is_prev_closed := q_data.ClosePosition
	var open_order_req *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	var close_order_req *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	if !is_prev_closed && prev_side != "" {
		if b.side.IsBuy(prev_side) && b.position_side.IsLong(prev_position_side) {
			if req.IsPositionSideLong() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide(b.side.Buy())
				open_order_req.SetType(b.order_type.Market())
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			} else if req.IsPositionSideShort() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide(b.side.Sell())
				open_order_req.SetType(b.order_type.Market())
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			}
			close_order_req = close_order_req.New()
			close_order_req.SetPositionSide(b.position_side.Long())
			close_order_req.SetSide(b.side.Sell())
			close_order_req.SetType(b.order_type.Market())
			close_order_req.SetEntryQuantity(q_data.ExecutedQty)
			close_order_req.SetSymbol(q_data.Symbol)
			close_order_req.SetClientOrderId(req.PrevClientId)
		} else if b.side.IsSell(prev_side) && b.position_side.IsShort(prev_position_side) {
			if req.IsPositionSideLong() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide(b.side.Buy())
				open_order_req.SetType(b.order_type.Market())
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			} else if req.IsPositionSideShort() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide(b.side.Sell())
				open_order_req.SetType(b.order_type.Market())
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			}
			close_order_req = close_order_req.New()
			close_order_req.SetPositionSide(b.position_side.Short())
			close_order_req.SetSide(b.side.Buy())
			close_order_req.SetType(b.order_type.Market())
			close_order_req.SetEntryQuantity(q_data.ExecutedQty)
			close_order_req.SetSymbol(q_data.Symbol)
			close_order_req.SetClientOrderId(req.PrevClientId)
		}
	} else {
		open_order_req = open_order_req.New()
		open_order_req.SetPositionSide(req.PositionSide)
		open_order_req.SetSide(b.side.Buy())
		open_order_req.SetType(b.order_type.Market())
		open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
		open_order_req.SetSymbol(req.Symbol)
		open_order_req.SetClientOrderId(req.CurrentClientId)
	}

	if close_order_req != nil {
		_, err := b.bn_service.PlaceSingleOrder(ctx, close_order_req)
		if err != nil {
			return nil, err
		}
	}
	if open_order_req != nil {
		if req.LeverageLevel != "" {
			_, err := b.bn_service.SetNewLeverage(ctx, &bnserivcemodelreq.SetLeverageBinanceServiceRequest{
				Symbol:   open_order_req.Symbol,
				Leverage: req.LeverageLevel,
			})
			if err != nil {
				return nil, err
			}
		}

		_, err = b.bn_service.PlaceSingleOrder(ctx, open_order_req)
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
