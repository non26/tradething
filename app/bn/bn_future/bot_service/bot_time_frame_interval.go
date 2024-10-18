package bnfuture

import (
	"context"
	"fmt"
	"strings"
	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"
)

func (t *botService) TimeIntervalSemiBotService(
	ctx context.Context,
	req *bothandlerreq.TradeTimeIntervalBinanceFutureRequest,
) (*bothandlerres.TradeTimeIntervalBinanceFutureResponse, error) {
	query_request := req.ToQueryOrderBinanceServiceRequest()
	q_data, err := t.bn_service.QueryOrder(
		ctx,
		query_request)
	if err != nil {
		return nil, err
	}

	prev_side := strings.ToUpper(q_data.Side)
	prev_position_side := strings.ToUpper(q_data.PositionSide)
	is_prev_closed := q_data.ClosePosition
	var open_order_req *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	var close_order_req *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	if !is_prev_closed && prev_side != "" {
		if prev_side == "BUY" && prev_position_side == "LONG" {
			if req.IsPositionSideLong() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide("BUY")
				open_order_req.SetType("MARKET")
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			} else if req.IsPositionSideShort() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide("SELL")
				open_order_req.SetType("MARKET")
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			}
			close_order_req = close_order_req.New()
			close_order_req.SetPositionSide("LONG")
			close_order_req.SetSide("SELL")
			close_order_req.SetType("MARKET")
			close_order_req.SetEntryQuantity(q_data.ExecutedQty)
			close_order_req.SetSymbol(q_data.Symbol)
			close_order_req.SetClientOrderId(req.PrevClientId)
		} else if prev_side == "SELL" && prev_position_side == "SHORT" {
			if req.IsPositionSideLong() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide("BUY")
				open_order_req.SetType("MARKET")
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			} else if req.IsPositionSideShort() {
				open_order_req = open_order_req.New()
				open_order_req.SetPositionSide(req.PositionSide)
				open_order_req.SetSide("SELL")
				open_order_req.SetType("MARKET")
				open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
				open_order_req.SetSymbol(req.Symbol)
				open_order_req.SetClientOrderId(req.CurrentClientId)
			}
			close_order_req = close_order_req.New()
			close_order_req.SetPositionSide("SHORT")
			close_order_req.SetSide("BUY")
			close_order_req.SetType("MARKET")
			close_order_req.SetEntryQuantity(q_data.ExecutedQty)
			close_order_req.SetSymbol(q_data.Symbol)
			close_order_req.SetClientOrderId(req.PrevClientId)
		}
	} else {
		open_order_req = open_order_req.New()
		open_order_req.SetPositionSide(req.PositionSide)
		open_order_req.SetSide("BUY")
		open_order_req.SetType("MARKET")
		open_order_req.SetEntryQuantity(fmt.Sprintf("%v", req.EntryQuantity))
		open_order_req.SetSymbol(req.Symbol)
		open_order_req.SetClientOrderId(req.CurrentClientId)
	}

	if close_order_req != nil {
		_, err := t.bn_service.PlaceSingleOrder(ctx, close_order_req)
		if err != nil {
			return nil, err
		}
	}
	if open_order_req != nil {
		if req.LeverageLevel != "" {
			_, err := t.bn_service.SetNewLeverage(ctx, &bnserivcemodelreq.SetLeverageBinanceServiceRequest{
				Symbol:   open_order_req.Symbol,
				Leverage: req.LeverageLevel,
			})
			if err != nil {
				return nil, err
			}
		}

		_, err = t.bn_service.PlaceSingleOrder(ctx, open_order_req)
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
