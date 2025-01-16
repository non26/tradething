package bnfuture

import ireqmodel "tradething/app/bn/bn_future/bnservice_request_model"

type CandleStickRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

func NewCandleStickRequest(c *CandleStickRequest) ireqmodel.IBnFutureRequest {
	return c
}

func (c *CandleStickRequest) PrepareRequest() {

}

func (c *CandleStickRequest) GetData() interface{} {
	return c
}
