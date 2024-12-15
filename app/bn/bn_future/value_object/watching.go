package bnfuture

type Watching struct {
	StopLoss   *StopLoss   `json:"stopLoss"`
	TakeProfit *TakeProfit `json:"takeProfit"`
}

type StopLoss struct {
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
	Interval string  `json:"interval"`
}

type TakeProfit struct {
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
	Interval string  `json:"interval"`
}
