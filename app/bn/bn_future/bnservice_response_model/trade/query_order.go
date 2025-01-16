package bnfuture

import (
	bnfuture "tradething/app/bn/bn_future/handler_response_model"
)

type QueryOrderBinanceServiceResponse struct {
	AvgPrice                string `json:"avgPrice"`
	ClientOrderID           string `json:"clientOrderId"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderID                 int    `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	OrigType                string `json:"origType"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	StopPrice               string `json:"stopPrice"`
	ClosePosition           bool   `json:"closePosition"`
	Symbol                  string `json:"symbol"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	PriceMatch              string `json:"priceMatch"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int    `json:"goodTillDate"`
}

func (q *QueryOrderBinanceServiceResponse) ToHandlerResponse() *bnfuture.QueryOrder {
	h := &bnfuture.QueryOrder{
		AvgPrice:                q.AvgPrice,
		ClientOrderID:           q.ClientOrderID,
		CumQuote:                q.CumQuote,
		ExecutedQty:             q.ExecutedQty,
		OrderID:                 q.OrderID,
		OrigQty:                 q.OrigQty,
		OrigType:                q.OrigType,
		Price:                   q.Price,
		ReduceOnly:              q.ReduceOnly,
		Side:                    q.Side,
		PositionSide:            q.PositionSide,
		Status:                  q.Status,
		StopPrice:               q.StopPrice,
		ClosePosition:           q.ClosePosition,
		Symbol:                  q.Symbol,
		Time:                    q.Time,
		TimeInForce:             q.TimeInForce,
		Type:                    q.Type,
		ActivatePrice:           q.ActivatePrice,
		PriceRate:               q.PriceRate,
		UpdateTime:              q.UpdateTime,
		WorkingType:             q.WorkingType,
		PriceProtect:            q.PriceProtect,
		PriceMatch:              q.PriceMatch,
		SelfTradePreventionMode: q.SelfTradePreventionMode,
		GoodTillDate:            q.GoodTillDate,
	}
	return h
}
