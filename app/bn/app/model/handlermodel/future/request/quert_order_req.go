package bnhandlerreq

import bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"

type QueryOrderBinanceHandlerRequest struct {
	Symbol string `json:"symbol"`
}

func (q *QueryOrderBinanceHandlerRequest) ToBinanceServiceQueryOrder() *bnserivcemodelreq.QueryOrderBinanceServiceRequest {
	m := bnserivcemodelreq.QueryOrderBinanceServiceRequest{
		Symbol: q.Symbol,
	}
	return &m
}
