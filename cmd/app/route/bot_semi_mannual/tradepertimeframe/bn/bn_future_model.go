package tradepertimeframe

type TradePerTimeFrameBinanceFutureRequest struct {
	Side          string  `json:"side"`
	PositionSide  string  `json:"positionSide"`
	EntryQuantity float64 `json:"entryQuantity"`
	Symbol        string  `json:"symbol"`
	LeverageLevel int16   `json:"leverageLevel"`
}
