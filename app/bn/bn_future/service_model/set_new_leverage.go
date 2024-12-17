package bnfuture

import (
	"strconv"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
)

type SetLeverageServiceRequest struct {
	leverage int
	symbol   string
}

func (s *SetLeverageServiceRequest) GetLeverage() int {
	return s.leverage
}

func (s *SetLeverageServiceRequest) SetLeverage(leverage int) {
	s.leverage = leverage
}

func (s *SetLeverageServiceRequest) GetSymbol() string {
	return s.symbol
}

func (s *SetLeverageServiceRequest) SetSymbol(symbol string) {
	s.symbol = symbol
}

func (s *SetLeverageServiceRequest) ToBinanceServiceSetLeverage() *bntradereq.SetLeverageBinanceServiceRequest {
	m := bntradereq.SetLeverageBinanceServiceRequest{
		Symbol:   s.symbol,
		Leverage: strconv.Itoa(s.leverage),
	}
	return &m
}
