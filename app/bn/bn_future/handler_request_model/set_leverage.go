package bnfuture

import (
	"strconv"
	"strings"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	model "tradething/app/bn/bn_future/service_model"
)

type SetLeverage struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

func (s *SetLeverage) Transform() {
	s.Symbol = strings.ToUpper(s.Symbol)
}

func (s *SetLeverage) ToBinanceServiceSetLeverage() *bntradereq.SetLeverage {
	m := bntradereq.SetLeverage{
		Symbol:   s.Symbol,
		Leverage: strconv.Itoa(s.Leverage),
	}
	return &m
}

func (s *SetLeverage) ToServiceModel() *model.Leverage {
	m := model.Leverage{}
	m.SetSymbol(s.Symbol)
	m.SetLeverage(s.Leverage)
	return &m
}

type SetLeverageBinanceHandlerResponse struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}
