package dto

import (
	"tradething/app/bn/future/BFF/market_data/domain"
	externalServiceDomain "tradething/app/bn/future/market_data/domain"
)

type Kline struct {
	Symbol           string
	Interval         string
	StartTime        int64
	EndTime          int64
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	CloseTime        int64
	QuoteVolume      string
	NumberOfTrades   int64
	BuyerBaseVolume  string
	BuyerQuoteVolume string
	IsGreenCandle    bool
}

func NewEmptyKline() *Kline {
	return &Kline{}
}

func (k *Kline) ToDomain() *domain.Kline {
	return &domain.Kline{
		Symbol:    k.Symbol,
		Interval:  k.Interval,
		StartTime: k.StartTime,
		EndTime:   k.EndTime,
		Open:      k.Open,
		High:      k.High,
		Low:       k.Low,
		Close:     k.Close,
	}
}

func (k *Kline) FromDomain(domainKline *domain.Kline) *Kline {
	return &Kline{
		Symbol:    domainKline.Symbol,
		Interval:  domainKline.Interval,
		StartTime: domainKline.StartTime,
		EndTime:   domainKline.EndTime,
		Open:      domainKline.Open,
		High:      domainKline.High,
		Low:       domainKline.Low,
		Close:     domainKline.Close,
	}
}

func (k *Kline) FromExternalServiceDomain(d *externalServiceDomain.Kline) *Kline {
	return &Kline{
		Symbol:        d.Symbol,
		Interval:      d.Interval,
		StartTime:     d.StartTime,
		EndTime:       d.EndTime,
		Open:          d.Open,
		High:          d.High,
		Low:           d.Low,
		Close:         d.Close,
		IsGreenCandle: d.IsGreenCandle,
	}
}

func (k *Kline) ToExternalServiceDomain() *externalServiceDomain.Kline {
	return &externalServiceDomain.Kline{
		Symbol:    k.Symbol,
		Interval:  k.Interval,
		StartTime: k.StartTime,
		EndTime:   k.EndTime,
		Open:      k.Open,
		High:      k.High,
		Low:       k.Low,
		Close:     k.Close,
	}
}
