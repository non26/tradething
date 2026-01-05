package res

import "fmt"

type NewOrderRes struct {
	// for error
	Code    *int    `json:"code"`
	Message *string `json:"msg"`
	// for success
	ClientOrderID           *string `json:"clientOrderId"`
	CumQty                  *string `json:"cumQty"`
	CumQuote                *string `json:"cumQuote"`
	ExecutedQty             *string `json:"executedQty"`
	OrderID                 *int    `json:"orderId"`
	AvgPrice                *string `json:"avgPrice"`
	OrigQty                 *string `json:"origQty"`
	Price                   *string `json:"price"`
	ReduceOnly              *bool   `json:"reduceOnly"`
	Side                    *string `json:"side"`
	PositionSide            *string `json:"positionSide"`
	Status                  *string `json:"status"`
	StopPrice               *string `json:"stopPrice"`
	ClosePosition           *bool   `json:"closePosition"`
	Symbol                  *string `json:"symbol"`
	TimeInForce             *string `json:"timeInForce"`
	Type                    *string `json:"type"`
	OrigType                *string `json:"origType"`
	ActivatePrice           *string `json:"activatePrice"`
	PriceRate               *string `json:"priceRate"`
	UpdateTime              *int64  `json:"updateTime"`
	WorkingType             *string `json:"workingType"`
	PriceProtect            *bool   `json:"priceProtect"`
	PriceMatch              *string `json:"priceMatch"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode"`
	GoodTillDate            *int64  `json:"goodTillDate"`
}

func (r *NewOrderRes) IsError() bool {
	return r.Code != nil && r.Message != nil
}

func (r *NewOrderRes) GetError() error {
	return fmt.Errorf("binance error code: %d, message: %s", *r.Code, *r.Message)
}
