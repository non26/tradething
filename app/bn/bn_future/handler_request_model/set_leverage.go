package bnfuture

import (
	"strconv"
	"strings"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

type SetLeverageHandlerRequest struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

func (s *SetLeverageHandlerRequest) Transform() {
	s.Symbol = strings.ToUpper(s.Symbol)
}

func (s *SetLeverageHandlerRequest) ToBinanceServiceSetLeverage() *bntradereq.SetLeverageBinanceServiceRequest {
	m := bntradereq.SetLeverageBinanceServiceRequest{
		Symbol:   s.Symbol,
		Leverage: strconv.Itoa(s.Leverage),
	}
	return &m
}

func (s *SetLeverageHandlerRequest) ToServiceModel() *svcfuture.SetLeverageServiceRequest {
	m := svcfuture.SetLeverageServiceRequest{}
	m.SetSymbol(s.Symbol)
	m.SetLeverage(s.Leverage)
	return &m
}

type SetLeverageBinanceHandlerResponse struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}
