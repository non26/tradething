package bnfuture

import (
	"fmt"
	"strings"
	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
	"tradething/app/bn/bncommon"
)

type TradeTimeIntervalBinanceFutureRequest struct {
	// Side          string  `json:"side"`
	PositionSide    string  `json:"positionSide"`  // long/short
	EntryQuantity   float64 `json:"entryQuantity"` // 0.005
	Symbol          string  `json:"symbol"`        // btcusdt
	PrevClientId    string  `json:"prevCliId"`
	CurrentClientId string  `json:"currCliId"`
	// LeverageLevel   string  `json:"leverageLevel"` // 125
}

func (t *TradeTimeIntervalBinanceFutureRequest) Validate() error {
	position_side := bncommon.NewPositionSide()
	if !(position_side.IsLong(t.PositionSide) || position_side.IsShort(t.PositionSide)) {
		return fmt.Errorf("invalid position side")
	}
	if t.Symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if t.EntryQuantity <= 0 {
		return fmt.Errorf("entry quantity must be greater than 0")
	}
	return nil
}

func (t *TradeTimeIntervalBinanceFutureRequest) ToUpper() {
	t.PositionSide = strings.ToUpper(t.PositionSide)
	t.Symbol = strings.ToUpper(t.Symbol)
}

func (t *TradeTimeIntervalBinanceFutureRequest) ToQueryOrderBinanceServiceRequest() *bnserivcemodelreq.QueryOrderBinanceServiceRequest {
	return &bnserivcemodelreq.QueryOrderBinanceServiceRequest{
		Symbol:            t.Symbol,
		OrigClientOrderId: t.PrevClientId,
	}
}
