package req

import bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"

type SpotOrderRequest struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	Type             string `json:"type"`
	Quantity         string `json:"quantity"`
	NewClientOrderId string `json:"newClientOrderId"`
	Timestamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}

// type IBnFutureServiceRequest interface {
// 	PrepareRequest()
// 	GetData() interface{}
// }

func (s *SpotOrderRequest) PrepareRequest() {}

func (s *SpotOrderRequest) GetData() interface{} {
	return s
}

func NewPlaceSignleOrderBinanceServiceRequest(
	s *SpotOrderRequest,
) bnrequest.IBnFutureServiceRequest {
	return s
}
