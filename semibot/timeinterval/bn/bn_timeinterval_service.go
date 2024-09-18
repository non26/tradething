package timeinterval

import (
	"context"
	"fmt"
	"strings"
	bnservice "tradething/app/bn/bn_future/bnservice"
	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
)

type bnTimeInterval struct {
	bn_service bnservice.IBinanceFutureExternalService
}

func NewBnTimeIntervalService(
	bn_service bnservice.IBinanceFutureExternalService,
) *bnTimeInterval {
	return &bnTimeInterval{
		bn_service,
	}
}

func (t *bnTimeInterval) TimeIntervalSemiBotService(
	ctx context.Context,
	req *TradeTimeIntervalBinanceFutureRequest,
) (*TradeTimeIntervalBinanceFutureResponse, error) {

	query_request := bnserivcemodelreq.QueryOrderBinanceServiceRequest{
		Symbol:            req.Symbol,
		OrigClientOrderId: req.PrevClientId,
	}
	q_data, err := t.bn_service.QueryOrder(
		ctx,
		&query_request)
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
			if req.PositionSide == "LONG" {
				open_order_req = CreateOpenOrder(
					req.PositionSide,
					"BUY",
					"MARKET",
					fmt.Sprintf("%v", req.EntryQuantity),
					req.Symbol,
					req.CurrentClientId,
				)
			} else if req.PositionSide == "SHORT" {
				open_order_req = CreateOpenOrder(
					req.PositionSide,
					"SELL",
					"MARKET",
					fmt.Sprintf("%v", req.EntryQuantity),
					req.Symbol,
					req.CurrentClientId,
				)
			}
			close_order_req = CreateOpenOrder(
				"LONG",
				"SELL",
				"MARKET",
				q_data.ExecutedQty,
				q_data.Symbol,
				req.PrevClientId,
			)
		} else if prev_side == "SELL" && prev_position_side == "SHORT" {
			if req.PositionSide == "LONG" {
				open_order_req = CreateOpenOrder(
					req.PositionSide,
					"BUY",
					"MARKET",
					fmt.Sprintf("%v", req.EntryQuantity),
					req.Symbol,
					req.CurrentClientId,
				)
			} else if req.PositionSide == "SHORT" {
				open_order_req = CreateOpenOrder(
					req.PositionSide,
					"SELL",
					"MARKET",
					fmt.Sprintf("%v", req.EntryQuantity),
					req.Symbol,
					req.CurrentClientId,
				)
			}
			close_order_req = CreateOpenOrder(
				"SHORT",
				"BUY",
				"MARKET",
				q_data.ExecutedQty,
				q_data.Symbol,
				req.PrevClientId,
			)
		}
	} else {
		open_order_req = CreateOpenOrder(
			req.PositionSide,
			"BUY",
			"MARKET",
			fmt.Sprintf("%v", req.EntryQuantity),
			req.Symbol,
			req.CurrentClientId,
		)
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
		return &TradeTimeIntervalBinanceFutureResponse{
			Message: "Suucess",
		}, nil
	}

	return &TradeTimeIntervalBinanceFutureResponse{
		Message: "Suucess But No Order Made",
	}, nil
}

func CreateOpenOrder(
	position_side string,
	side string,
	order_type string,
	quantity string,
	symbol string,
	clientId string,

) *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest {
	order := &bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest{
		PositionSide:  position_side,
		Side:          side,
		Type:          order_type,
		EntryQuantity: quantity,
		Symbol:        symbol,
		ClientOrderId: clientId,
	}
	return order
}
