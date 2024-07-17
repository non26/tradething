package bnserivcemodelreq

import (
	"strconv"
	"strings"
	"tradething/app/bn/bncommon"
)

type QueryOrderBinanceServiceRequest struct {
	Symbol string `json:"symbol"`
	// OrderId           float64 `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
	// RecvWindow        float64 `json:"recvWindow"`
	Timestamp string `json:"timestamp"`
}

func (q *QueryOrderBinanceServiceRequest) PrepareRequest() {
	q.Symbol = strings.ToUpper(q.Symbol)
	q.setTimeStamp()
	q.checkClientId()
}

func (q *QueryOrderBinanceServiceRequest) GetData() interface{} {
	return q
}

func (q *QueryOrderBinanceServiceRequest) setTimeStamp() {
	q.Timestamp = strconv.FormatInt(bncommon.GetTimeStamp(), 10)
}

func (q *QueryOrderBinanceServiceRequest) checkClientId() {
	if q.OrigClientOrderId == "" {
		q.OrigClientOrderId = q.Symbol
	}
}

func NewQueryOrderBinanceServiceRequest(
	q *QueryOrderBinanceServiceRequest,
) IBnFutureServiceRequest {
	return q
}
