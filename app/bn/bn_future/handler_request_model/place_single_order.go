package bnfuture

import (
	svcfuture "tradething/app/bn/bn_future/service_model"
)

type PlaceSignleOrderHandlerRequest struct {
	PositionSide  string `json:"positionSide"`
	Side          string `json:"side"`
	EntryQuantity string `json:"entryQuantity"`
	Symbol        string `json:"symbol"`
	LeverageLevel int    `json:"leverageLevel"`
	ClientOrderId string `json:"newClientOrderId"`
}

func (p *PlaceSignleOrderHandlerRequest) ToServiceModel() *svcfuture.PlaceSignleOrderServiceRequest {
	m := svcfuture.PlaceSignleOrderServiceRequest{}
	m.SetPositionSide(p.PositionSide)
	m.SetSide(p.Side)
	m.SetEntryQuantity(p.EntryQuantity)
	m.SetSymbol(p.Symbol)
	m.SetClientOrderId(p.ClientOrderId)
	return &m
}

type PlaceSignleOrderHandlerResponse struct {
	Symbol   string `json:"symbol"`
	Quantity string `json:"quantity"`
}
