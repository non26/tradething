package bnfuture

import bnfuture "tradething/app/bn/bn_future/bnservice_request_model"

type QueryOrderBinanceHandlerRequest struct {
	Symbol            string `json:"symbol"`
	OrigClientOrderId string `json:"origClientOrderId"`
}

func (q *QueryOrderBinanceHandlerRequest) ToBinanceServiceQueryOrder() *bnfuture.QueryOrderBinanceServiceRequest {
	m := bnfuture.QueryOrderBinanceServiceRequest{
		Symbol:            q.Symbol,
		OrigClientOrderId: q.OrigClientOrderId,
	}
	return &m
}

type QueryOrderBinanceHandlerResponse struct {
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
