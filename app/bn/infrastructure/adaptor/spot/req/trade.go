package req

import (
	"strconv"

	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

type SpotOrderRequest struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	Type             string `json:"type"`
	Quantity         string `json:"quantity"`
	NewClientOrderId string `json:"newClientOrderId"`
	Timestamp        string `json:"timestamp"`
}

// type IBnFutureServiceRequest interface {
// 	PrepareRequest()
// 	GetData() interface{}
// }

func (s *SpotOrderRequest) PrepareRequest() {
	s.setTimestamp()
}

func (s *SpotOrderRequest) GetData() interface{} {
	return s
}

func (s *SpotOrderRequest) setTimestamp() {
	s.Timestamp = strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
}

func NewPlaceSignleOrderBinanceServiceRequest(
	s *SpotOrderRequest,
) bnrequest.IBnFutureServiceRequest {
	return s
}
