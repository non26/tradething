package bnfuture

import (
	"strconv"
	bnfuture "tradething/app/bn/bn_future/bnservice_request_model"
)

type SetLeverageHandlerRequest struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

func (s *SetLeverageHandlerRequest) ToBinanceServiceSetLeverage() *bnfuture.SetLeverageBinanceServiceRequest {
	m := bnfuture.SetLeverageBinanceServiceRequest{
		Symbol:   s.Symbol,
		Leverage: strconv.Itoa(s.Leverage),
	}
	return &m
}

type SetLeverageBinanceHandlerResponse struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}
