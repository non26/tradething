package bnfuture

import bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"

type QueryOrderServiceRequest struct {
	symbol            string
	origClientOrderId string
}

func (q *QueryOrderServiceRequest) GetSymbol() string {
	return q.symbol
}

func (q *QueryOrderServiceRequest) SetSymbol(symbol string) {
	q.symbol = symbol
}

func (q *QueryOrderServiceRequest) GetOrigClientOrderId() string {
	return q.origClientOrderId
}

func (q *QueryOrderServiceRequest) SetOrigClientOrderId(origClientOrderId string) {
	q.origClientOrderId = origClientOrderId
}

func (q *QueryOrderServiceRequest) ToBinanceServiceQueryOrder() *bntradereq.QueryOrderBinanceServiceRequest {
	m := bntradereq.QueryOrderBinanceServiceRequest{
		Symbol:            q.symbol,
		OrigClientOrderId: q.origClientOrderId,
	}
	return &m
}
