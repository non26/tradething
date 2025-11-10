package res

import "tradething/app/bn/future/position/domain"

type GetAllCurrentPositionRes struct {
	Positions []*GetCurrentPositionRes `json:"positions"`
}

func (r *GetAllCurrentPositionRes) FromDomain(positions []*domain.Position) *GetAllCurrentPositionRes {
	for _, position := range positions {
		r.Positions = append(r.Positions, &GetCurrentPositionRes{
			ClientId:     position.ClientId,
			Symbol:       position.Symbol,
			PositionSide: position.PositionSide,
			Side:         position.Side,
			AmountB:      position.AmountB,
			AccountId:    position.AccountId,
		})
	}
	return r
}
