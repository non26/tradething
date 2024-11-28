package bnfuture

import (
	"strconv"
	"strings"

	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

type SetLeverageBinanceServiceRequest struct {
	Leverage  string `json:"leverage"`
	Symbol    string `json:"symbol"`
	Timestamp string `json:"timestamp"`
}

func (s *SetLeverageBinanceServiceRequest) PrepareRequest() {
	s.Symbol = strings.ToUpper(s.Symbol)
	s.setTimeStamp()
}

func (s *SetLeverageBinanceServiceRequest) GetData() interface{} {
	return s
}

func (s *SetLeverageBinanceServiceRequest) setTimeStamp() {
	s.Timestamp = strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
}

func NewSetLeverageBinanceServiceRequest(
	s *SetLeverageBinanceServiceRequest,
) IBnFutureServiceRequest {
	return s
}
