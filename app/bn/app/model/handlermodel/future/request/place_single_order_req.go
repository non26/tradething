package bnhandlerreq

import (
	"fmt"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
)

type PlaceSignleOrderHandlerRequest struct {
	PositionSide  string  `json:"positionSide"`
	Side          string  `json:"side"`
	EntryQuantity float64 `json:"entryQuantity"`
	Symbol        string  `json:"symbol"`
	LeverageLevel int     `json:"leverageLevel"`
	ClientOrderId string  `json:"newClientOrderId"`
}

func (p *PlaceSignleOrderHandlerRequest) ToBinanceServiceModel() *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest {
	m := bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest{
		PositionSide:  p.PositionSide,
		Side:          p.Side,
		EntryQuantity: fmt.Sprintf("%v", p.EntryQuantity),
		Symbol:        p.Symbol,
		ClientOrderId: p.ClientOrderId,
	}
	return &m
}
