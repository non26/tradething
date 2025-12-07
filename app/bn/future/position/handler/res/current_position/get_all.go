package res

import "tradething/app/bn/future/position/domain"

type GetAllCurrentPositionRes struct {
	Positions []*GetCurrentPositionRes `json:"positions"`
}

func (r *GetAllCurrentPositionRes) FromDomain(positions []*domain.Position) *GetAllCurrentPositionRes {
	if len(positions) == 0 {
		return nil
	}
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
