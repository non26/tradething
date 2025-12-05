package res

import "tradething/app/bn/future/accumulation/domain"

type GetRes struct {
	AccumID      string `json:"accum_id"`
	ClientID     string `json:"client_id"`
	MaxAccum     string `json:"max_accum"`
	PresentAccum string `json:"present_accum"`
}

func (r *GetRes) FromDomain(accumulation *domain.Accumulation) *GetRes {
	if accumulation == nil {
		return nil
	}
	r.AccumID = accumulation.AccumID
	r.ClientID = accumulation.ClientID
	r.MaxAccum = accumulation.MaxAccum
	r.PresentAccum = accumulation.PresentAccum
	return r
}
