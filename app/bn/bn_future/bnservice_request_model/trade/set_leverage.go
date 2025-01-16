package bnfuture

import (
	"strconv"
	"strings"

	ireqmodel "tradething/app/bn/bn_future/bnservice_request_model"

	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

type SetLeverage struct {
	Leverage  string `json:"leverage"`
	Symbol    string `json:"symbol"`
	Timestamp string `json:"timestamp"`
}

func (s *SetLeverage) PrepareRequest() {
	s.Symbol = strings.ToUpper(s.Symbol)
	s.setTimeStamp()
}

func (s *SetLeverage) GetData() interface{} {
	return s
}

func (s *SetLeverage) setTimeStamp() {
	s.Timestamp = strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
}

func NewSetLeverageBinanceServiceRequest(
	s *SetLeverage,
) ireqmodel.IBnFutureRequest {
	return s
}
