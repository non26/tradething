package bnfuture

import bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"

type Order struct {
	symbol            string
	origClientOrderId string
}

func (q *Order) GetSymbol() string {
	return q.symbol
}

func (q *Order) SetSymbol(symbol string) {
	q.symbol = symbol
}

func (q *Order) GetOrigClientOrderId() string {
	return q.origClientOrderId
}

func (q *Order) SetOrigClientOrderId(origClientOrderId string) {
	q.origClientOrderId = origClientOrderId
}

func (q *Order) ToBinanceServiceQueryOrder() *bntradereq.QueryOrderBinanceServiceRequest {
	m := bntradereq.QueryOrderBinanceServiceRequest{
		Symbol:            q.symbol,
		OrigClientOrderId: q.origClientOrderId,
	}
	return &m
}
