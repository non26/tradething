package res

import "tradething/app/bn/future/position/domain"

type GetAllAdvancedPositionRes struct {
	Positions []*GetAdvancedPositionRes `json:"positions"`
}

func (r *GetAllAdvancedPositionRes) FromDomain(positions []*domain.Position) *GetAllAdvancedPositionRes {
	for _, position := range positions {
		r.Positions = append(r.Positions, &GetAdvancedPositionRes{
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
