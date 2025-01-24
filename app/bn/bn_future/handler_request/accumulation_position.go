package bnfuture

type AccumulatePosition struct {
	AccumulateId     string           `json:"accummulateId"`
	Symbol           string           `json:"symbol"`
	PositionSide     string           `json:"position_side"`
	Side             string           `json:"side"`
	AccumulateConfig AccumulateConfig `json:"accumulate_config"`
}

type AccumulateConfig struct {
	UpperBoundPrice     float64 `json:"upper_bound_price"`
	LowerBoundPrice     float64 `json:"lower_bound_price"`
	TotalAmountQ        string  `json:"total_amount_q"`
	AccumulationAmountQ string  `json:"accumulation_amount_q"`
}
