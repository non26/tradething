package bnserivcemodelreq

import (
	"strconv"
	"strings"
	"tradething/app/bn/bncommon"
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

func (s *SetLeverageBinanceServiceRequest) setTimeStamp() {
	s.Timestamp = strconv.FormatInt(bncommon.GetTimeStamp(), 10)
}
