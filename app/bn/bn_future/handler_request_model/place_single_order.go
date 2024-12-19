package bnfuture

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	svcfuture "tradething/app/bn/bn_future/service_model"
	valueobject "tradething/app/bn/bn_future/value_object"
)

var minute_interval = []string{"30m", "45m"}
var hour_interval = []string{"1h", "2h", "3h", "4h", "6h", "8h", "12h"}

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

	if strings.Contains(p.Watching.StopLoss.Interval, "m") {
		if !slices.Contains(minute_interval, p.Watching.StopLoss.Interval) {
			return errors.New("invalid minute interval")
		}
	} else if strings.Contains(p.Watching.StopLoss.Interval, "h") {
		if !slices.Contains(hour_interval, p.Watching.StopLoss.Interval) {
			return errors.New("invalid hour interval")
		}
	} else if strings.Contains(p.Watching.StopLoss.Interval, "d") {
		period := p.Watching.StopLoss.Interval[:len(p.Watching.StopLoss.Interval)-1]
		_, err := strconv.Atoi(period)
		if err != nil {
			return errors.New("invalid day interval")
		}
	} else {
		return errors.New("invalid interval")
	}

	return nil
}

func (p *PlaceSignleOrderHandlerRequest) Transform() {
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.Side = strings.ToUpper(p.Side)
	p.Watching.StopLoss.Interval = strings.ToLower(p.Watching.StopLoss.Interval)

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
