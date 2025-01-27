package bnfuture

import (
	"fmt"
	"strings"
	model "tradething/app/bn/bn_future/service_model"
)

var minute_interval = []string{"30m", "45m"}
var hour_interval = []string{"1h", "2h", "3h", "4h", "6h", "8h", "12h"}

type PlacePosition struct {
	PositionSide string  `json:"positionSide"`
	Side         string  `json:"side"`
	AmountB      float64 `json:"amount_b"`
	Symbol       string  `json:"symbol"`
	ClientId     string  `json:"client_id"`
}

func (p *PlacePosition) Transform() {
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.Side = strings.ToUpper(p.Side)
}

func (p *PlacePosition) ToServiceModel() *model.Position {
	m := model.Position{}
	m.SetPositionSide(p.PositionSide)
	m.SetSide(p.Side)
	m.SetEntryQuantity(fmt.Sprintf("%f", p.AmountB))
	m.SetSymbol(p.Symbol)
	m.SetClientId(p.ClientId)
	return &m
}
