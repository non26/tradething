package bnfuture

import (
	bnfuture "tradething/app/bn/bn_future/handler_response_model"
)

type SetLeverageBinanceServiceResponse struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

func (s *SetLeverageBinanceServiceResponse) ToHandlerResponse() *bnfuture.SetLeverageBinanceHandlerResponse {
	m := bnfuture.SetLeverageBinanceHandlerResponse{
		Leverage: s.Leverage,
		Symbol:   s.Symbol,
	}
	return &m
}
