package bnfuture

import (
	"strconv"
	bntradereq "tradething/app/bn/bn_future/bnservice_request/trade"
)

type Leverage struct {
	leverage int
	symbol   string
}

func (s *Leverage) GetLeverage() int {
	return s.leverage
}

func (s *Leverage) SetLeverage(leverage int) {
	s.leverage = leverage
}

func (s *Leverage) GetSymbol() string {
	return s.symbol
}

func (s *Leverage) SetSymbol(symbol string) {
	s.symbol = symbol
}

func (s *Leverage) ToBinanceServiceSetLeverage() *bntradereq.SetLeverage {
	m := bntradereq.SetLeverage{
		Symbol:   s.symbol,
		Leverage: strconv.Itoa(s.leverage),
	}
	return &m
}
