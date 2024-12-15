package bnfuture

import (
	"errors"
	"fmt"
	"strings"
	svcfuture "tradething/app/bn/bn_future/service_model"
	valueobject "tradething/app/bn/bn_future/value_object"
)

type PlaceSignleOrderHandlerRequest struct {
	PositionSide  string                `json:"positionSide"`
	Side          string                `json:"side"`
	EntryQuantity float64               `json:"entryQuantity"`
	Symbol        string                `json:"symbol"`
	LeverageLevel int                   `json:"leverageLevel"`
	ClientOrderId string                `json:"newClientOrderId"`
	CurrentPrice  float64               `json:"currentPrice"`
	Watching      *valueobject.Watching `json:"watching"`
}

func (p *PlaceSignleOrderHandlerRequest) Validate() error {
	if p.Watching.StopLoss == nil {
		return errors.New("stopLoss is required")
	}
	return nil
}

func (p *PlaceSignleOrderHandlerRequest) Transform() {
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.Side = strings.ToUpper(p.Side)
	p.Watching.StopLoss.Interval = strings.ToLower(p.Watching.StopLoss.Interval)
	p.Watching.TakeProfit.Interval = strings.ToLower(p.Watching.TakeProfit.Interval)
}

func (p *PlaceSignleOrderHandlerRequest) ToServiceModel() *svcfuture.PlaceSignleOrderServiceRequest {
	m := svcfuture.PlaceSignleOrderServiceRequest{}
	m.SetPositionSide(p.PositionSide)
	m.SetSide(p.Side)
	m.SetEntryQuantity(fmt.Sprintf("%f", p.EntryQuantity))
	m.SetSymbol(p.Symbol)
	m.SetLeverageLevel(p.LeverageLevel)
	m.SetClientOrderId(p.ClientOrderId)
	m.SetLeverageLevel(p.LeverageLevel)
	m.SetStopLoss(p.Watching.StopLoss)
	m.SetTakeProfit(p.Watching.TakeProfit)
	return &m
}
