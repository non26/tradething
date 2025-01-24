package bnfuture

import valueobject "tradething/app/bn/bn_future/value_object"

type AdvancedPosition struct {
	PositionSide  string                `json:"positionSide"`
	Side          string                `json:"side"`
	EntryQuantity float64               `json:"entryQuantity"`
	Symbol        string                `json:"symbol"`
	LeverageLevel int                   `json:"leverageLevel"`
	ClientOrderId string                `json:"newClientOrderId"`
	CurrentPrice  float64               `json:"currentPrice"`
	Watching      *valueobject.Watching `json:"watching"`
}
