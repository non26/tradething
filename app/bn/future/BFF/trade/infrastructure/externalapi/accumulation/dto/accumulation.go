package dto

import (
	"tradething/app/bn/future/BFF/trade/domain"
	accumulationDomain "tradething/app/bn/future/accumulation/domain"
)

type AccumulationDto struct {
	AccumID      string
	ClientID     string
	MaxAccum     string
	PresentAccum string
}

func NewAccumulationDto() *AccumulationDto {
	return &AccumulationDto{}
}

func (a *AccumulationDto) FromDomain(d *domain.Order) *accumulationDomain.Accumulation {
	return &accumulationDomain.Accumulation{
		AccumID:      d.AccumID,
		ClientID:     d.ClientId,
		MaxAccum:     d.MaxAccum,
		PresentAccum: d.AmountB,
	}
}

func (a *AccumulationDto) ToDomain(d *accumulationDomain.Accumulation) *domain.Order {
	return &domain.Order{
		AccumID:  d.AccumID,
		ClientId: d.ClientID,
		MaxAccum: d.MaxAccum,
		AmountB:  d.PresentAccum,
	}
}
