package bnfuture

import "strings"

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
