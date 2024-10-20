package bncommon

import "strings"

type ISide interface {
	Buy() string
	Sell() string
	IsBuy(side string) bool
	IsSell(side string) bool
}

type side struct {
	buy  string
	sell string
}

func (s *side) Buy() string {
	return s.buy
}

func (s *side) Sell() string {
	return s.sell
}

func (s *side) IsBuy(side string) bool {
	return s.buy == s.sideTrasform(side)
}

func (s *side) IsSell(side string) bool {
	return s.sell == s.sideTrasform(side)
}

func (s *side) sideTrasform(side string) string {
	return strings.ToUpper(side)
}

func NewSide() ISide {
	return &side{
		buy:  "BUY",
		sell: "SELL",
	}
}
