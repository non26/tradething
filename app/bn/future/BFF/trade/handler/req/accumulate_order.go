package req

import "tradething/app/bn/future/BFF/trade/domain"

type AccumulateOrderReq struct {
	AccumID           string `json:"accum_id"`
	ReferenceClientID string `json:"reference_client_id"`
	MaxAccum          string `json:"max_accum"`
	Accum             string `json:"accum"`
}

func (r *AccumulateOrderReq) ToDomain() *domain.Order {
	return &domain.Order{
		AccumID:  r.AccumID,
		ClientId: r.ReferenceClientID,
		MaxAccum: r.MaxAccum,
		Accum:    r.Accum,
	}
}
