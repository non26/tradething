package res

import "tradething/app/bn/future/BFF/advanced_position/domain"

type GetAllRes struct {
	Positions []*GetRes `json:"positions"`
}

func NewGetAllRes() *GetAllRes {
	return &GetAllRes{
		Positions: make([]*GetRes, 0),
	}
}

func (r *GetAllRes) FromDomain(positions []*domain.AdvancedPosition) *GetAllRes {
	if len(positions) == 0 {
		return nil
	}
	for _, position := range positions {
		singlePosition := &GetRes{}
		r.Positions = append(r.Positions, singlePosition.FromDomain(position))
	}
	return r
}
