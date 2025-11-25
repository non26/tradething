package res

import "tradething/app/bn/future/BFF/advanced_position/domain"

type GetAllRes struct {
	Positions []*GetRes `json:"positions"`
}

func (r *GetAllRes) FromDomain(positions []*domain.AdvancedPosition) *GetAllRes {
	for _, position := range positions {
		singlePosition := &GetRes{}
		r.Positions = append(r.Positions, singlePosition.FromDomain(position))
	}
	return r
}
