package bnservicemodelres

import bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"

type PlaceSignleOrderBinanceServiceResponse struct {
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
}

func (p *PlaceSignleOrderBinanceServiceResponse) ToBnHandlerResponse() *bnhandlerres.PlaceSignleOrderHandlerResponse {
	m := bnhandlerres.PlaceSignleOrderHandlerResponse{
		Symbol:   p.Symbol,
		Quantity: p.Quantity,
	}
	return &m
}
