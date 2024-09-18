package bnfuture

import (
	bnfuture "tradething/app/bn/bn_future/handler_response_model"
)

type PlaceSignleOrderBinanceServiceResponse struct {
	Symbol   string `json:"symbol"`
	Quantity string `json:"origQty"`
}

func (p *PlaceSignleOrderBinanceServiceResponse) ToBnHandlerResponse() *bnfuture.PlaceSignleOrderHandlerResponse {
	m := bnfuture.PlaceSignleOrderHandlerResponse{
		Symbol:   p.Symbol,
		Quantity: p.Quantity,
	}
	return &m
}
