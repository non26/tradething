package req

import (
	"errors"
	"strings"
	"tradething/app/bn/process/future/domain"
)

type SetAdvancedPositionReqs struct {
	Positions []SetAdvancedPositionReq `json:"positions"`
}

func (s *SetAdvancedPositionReqs) ToDomain() []*domain.Position {
	positions := []*domain.Position{}
	for _, position := range s.Positions {
		positions = append(positions, position.ToDomain())
	}
	return positions
}

func (s *SetAdvancedPositionReqs) Validate() error {
	for _, position := range s.Positions {
		if err := position.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetAdvancedPositionReq struct {
	ClientId     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	PositionSide string `json:"position_side"`
	AmountB      string `json:"amount_b"`
}

func (s *SetAdvancedPositionReq) Transform() {
	s.Symbol = strings.ToUpper(s.Symbol)
	s.Side = strings.ToUpper(s.Side)
	s.PositionSide = strings.ToUpper(s.PositionSide)
}

func (s *SetAdvancedPositionReq) Validate() error {
	if s.ClientId == "" {
		return errors.New("client_id is required")
	}
	if s.Symbol == "" {
		return errors.New("symbol is required")
	}
	if s.Side == "" {
		return errors.New("side is required")
	}
	if s.PositionSide == "" {
		return errors.New("position_side is required")
	}
	if s.AmountB == "" {
		return errors.New("amount_b is required")
	}
	return nil
}

func (s *SetAdvancedPositionReq) ToDomain() *domain.Position {
	position := &domain.Position{}
	position.SetClientId(s.ClientId)
	position.SetSymbol(s.Symbol)
	position.SetSide(s.Side)
	position.SetPositionSide(s.PositionSide)
	position.SetEntryQuantity(s.AmountB)
	return position
}
