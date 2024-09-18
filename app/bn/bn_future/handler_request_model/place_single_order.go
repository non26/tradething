package bnfuture

import (
	"fmt"
	bnfuture "tradething/app/bn/bn_future/bnservice_request_model"
)

type PlaceSignleOrderHandlerRequest struct {
	PositionSide  string  `json:"positionSide"`
	Side          string  `json:"side"`
	EntryQuantity float64 `json:"entryQuantity"`
	Symbol        string  `json:"symbol"`
	LeverageLevel int     `json:"leverageLevel"`
	ClientOrderId string  `json:"newClientOrderId"`
}

func (p *PlaceSignleOrderHandlerRequest) ToBinanceServiceModel() *bnfuture.PlaceSignleOrderBinanceServiceRequest {
	m := bnfuture.PlaceSignleOrderBinanceServiceRequest{
		PositionSide:  p.PositionSide,
		Side:          p.Side,
		EntryQuantity: fmt.Sprintf("%v", p.EntryQuantity),
		Symbol:        p.Symbol,
		ClientOrderId: p.ClientOrderId,
	}
	return &m
}

type PlaceSignleOrderHandlerResponse struct {
	Symbol   string `json:"symbol"`
	Quantity string `json:"quantity"`
}
