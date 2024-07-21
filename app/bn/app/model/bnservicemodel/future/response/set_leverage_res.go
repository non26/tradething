package bnservicemodelres

import bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"

type SetLeverageBinanceServiceResponse struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

func (s *SetLeverageBinanceServiceResponse) ToHandlerResponse() *bnhandlerres.SetLeverageBinanceHandlerResponse {
	m := bnhandlerres.SetLeverageBinanceHandlerResponse{
		Leverage: s.Leverage,
		Symbol:   s.Symbol,
	}
	return &m
}
