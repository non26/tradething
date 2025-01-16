package bnfuture

import (
	bnfuture "tradething/app/bn/bn_future/handler_response"
)

type PlacePositionData struct {
	Symbol   string `json:"symbol"`
	Quantity string `json:"origQty"`
}

func (p *PlacePositionData) ToBnHandlerResponse() *bnfuture.PlacePosition {
	m := bnfuture.PlacePosition{
		Symbol:   p.Symbol,
		Quantity: p.Quantity,
	}
	return &m
}
