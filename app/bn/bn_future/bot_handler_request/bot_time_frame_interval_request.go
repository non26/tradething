package bnfuture

import (
	"strings"
	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
)

type TradeTimeIntervalBinanceFutureRequest struct {
	// Side          string  `json:"side"`
	PositionSide    string  `json:"positionSide"`  // long/short
	EntryQuantity   float64 `json:"entryQuantity"` // 0.005
	Symbol          string  `json:"symbol"`        // btcusdt
	PrevClientId    string  `json:"prevCliId"`
	CurrentClientId string  `json:"currCliId"`
	LeverageLevel   string  `json:"leverageLevel"` // 125
}

func (t *TradeTimeIntervalBinanceFutureRequest) ToUpper() {
	t.PositionSide = strings.ToUpper(t.PositionSide)
	t.Symbol = strings.ToUpper(t.Symbol)
}

func (t TradeTimeIntervalBinanceFutureRequest) IsPositionSideLong() bool {
	return t.PositionSide == "LONG"
}

func (t TradeTimeIntervalBinanceFutureRequest) IsPositionSideShort() bool {
	return t.PositionSide == "SHORT"
}

func (t *TradeTimeIntervalBinanceFutureRequest) ToQueryOrderBinanceServiceRequest() *bnserivcemodelreq.QueryOrderBinanceServiceRequest {
	return &bnserivcemodelreq.QueryOrderBinanceServiceRequest{
		Symbol:            t.Symbol,
		OrigClientOrderId: t.PrevClientId,
	}
}
