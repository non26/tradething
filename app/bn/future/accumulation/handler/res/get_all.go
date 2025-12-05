package res

import "tradething/app/bn/future/accumulation/domain"

type GetAllRes struct {
	Accumulations []*GetRes `json:"accumulations"`
}

func (r *GetAllRes) FromDomain(accumulations []*domain.Accumulation) *GetAllRes {
	if len(accumulations) == 0 {
		return nil
	}
	r.Accumulations = make([]*GetRes, len(accumulations))
	for i, accumulation := range accumulations {
		r.Accumulations[i] = &GetRes{}
		r.Accumulations[i].FromDomain(accumulation)
	}
	return r
}
