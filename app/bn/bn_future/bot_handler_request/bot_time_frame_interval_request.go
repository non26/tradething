package bnfuture

import (
	"fmt"
	"strings"
	svcrepo "tradething/app/bn/bn_future/repository_model"
	"tradething/app/bn/bncommon"
)

type TradeTimeIntervalBinanceFutureRequest struct {
	// Side          string  `json:"side"`
	PositionSide  string  `json:"positionSide"`  // long/short
	EntryQuantity float64 `json:"entryQuantity"` // 0.005
	Symbol        string  `json:"symbol"`        // btcusdt
	// PrevClientId    string  `json:"prevCliId"`
	// CurrentClientId string  `json:"currCliId"`
	LeverageLevel string `json:"leverageLevel"` // 125
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

func (t *TradeTimeIntervalBinanceFutureRequest) ToBnFutureOpeningPositionEntity(side string, leverage string, clientId string) *svcrepo.BinanceFutureOpeningPosition {
	return &svcrepo.BinanceFutureOpeningPosition{
		Symbol:             t.Symbol,
		PositionSide:       t.PositionSide,
		ClientId:           clientId,
		AmountQ:            fmt.Sprintf("%v", t.EntryQuantity),
		Side:               side,
		ExchangeId:         0,
		Leverage:           "",
		AmountB:            "",
		BuyOrderCreatedAt:  "",
		SellOrderCreatedAt: "",
	}
}

func (t *TradeTimeIntervalBinanceFutureRequest) ToBnFutureQouteUSDTEntity(countingSymbol int) *svcrepo.BinanceFutureQouteUSDT {
	return &svcrepo.BinanceFutureQouteUSDT{
		Symbol:         t.Symbol,
		CountingSymbol: countingSymbol,
	}
}
