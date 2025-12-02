package req

import "tradething/app/bn/future/accumulation/domain"

type UpsertReq struct {
	AccumID      string `json:"accum_id"`
	ClientID     string `json:"client_id"`
	MaxAccum     string `json:"max_accum"`
	PresentAccum string `json:"present_accum"`
}

func (r *UpsertReq) ToDomain() *domain.Accumulation {
	return &domain.Accumulation{
		AccumID:      r.AccumID,
		ClientID:     r.ClientID,
		MaxAccum:     r.MaxAccum,
		PresentAccum: r.PresentAccum,
	}
}
