package bnhandlerreq

import (
	"strconv"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
)

type SetLeverageHandlerRequest struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

func (s *SetLeverageHandlerRequest) ToBinanceServiceSetLeverage() *bnserivcemodelreq.SetLeverageBinanceServiceRequest {
	m := bnserivcemodelreq.SetLeverageBinanceServiceRequest{
		Symbol:   s.Symbol,
		Leverage: strconv.Itoa(s.Leverage),
	}
	return &m
}
