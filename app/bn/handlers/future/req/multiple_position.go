package req

import "tradething/app/bn/process/future/domain"

type MultiplePosition struct {
	Positions []Position `json:"positions"`
}

func (m *MultiplePosition) ToDomain() []*domain.Position {
	positions := []*domain.Position{}
	for _, position := range m.Positions {
		positions = append(positions, position.ToDomain())
	}
	return positions
}
