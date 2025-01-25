package bnfuture

import (
	model "tradething/app/bn/bn_future/service_model"

	"github.com/shopspring/decimal"
)

type AdvancedPosition struct {
	PositionSide  string  `json:"position_side"`
	Side          string  `json:"side"`
	EntryQuantity float64 `json:"entry_quantity"`
	Symbol        string  `json:"symbol"`
	LeverageLevel int     `json:"leverage_level"`
	ClientId      string  `json:"new_client_id"`
}

func (a *AdvancedPosition) ToPosition() *model.Position {
	m := new(model.Position)
	m.SetSymbol(a.Symbol)
	m.SetSide(a.Side)
	m.SetPositionSide(a.PositionSide)
	m.SetEntryQuantity(decimal.NewFromFloat(a.EntryQuantity).String())
	m.SetLeverageLevel(a.LeverageLevel)
	m.SetClientId(a.ClientId)
	return m
}
