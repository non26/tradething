package bnfuture

import (
	"fmt"
	"strings"
	svcfuture "tradething/app/bn/bn_future/service_model"
)

type PlaceSignleOrderHandlerRequest struct {
	PositionSide  string  `json:"positionSide"`
	Side          string  `json:"side"`
	EntryQuantity float64 `json:"entryQuantity"`
	Symbol        string  `json:"symbol"`
	LeverageLevel int     `json:"leverageLevel"`
	ClientOrderId string  `json:"newClientOrderId"`
}

func (p *PlaceSignleOrderHandlerRequest) Transform() {
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.Side = strings.ToUpper(p.Side)
}

func (p *PlaceSignleOrderHandlerRequest) ToServiceModel() *svcfuture.PlaceSignleOrderServiceRequest {
	m := svcfuture.PlaceSignleOrderServiceRequest{}
	m.SetPositionSide(p.PositionSide)
	m.SetSide(p.Side)
	m.SetEntryQuantity(fmt.Sprintf("%f", p.EntryQuantity))
	m.SetSymbol(p.Symbol)
	m.SetClientOrderId(p.ClientOrderId)
	return &m
}

type PlaceSignleOrderHandlerResponse struct {
	Symbol   string `json:"symbol"`
	Quantity string `json:"quantity"`
}
