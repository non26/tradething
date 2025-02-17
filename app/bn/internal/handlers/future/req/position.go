package req

import (
	"errors"
	"strings"
	"tradething/app/bn/internal/process/future/domain"
)

type Position struct {
	PositionSide string `json:"positionSide"`
	Side         string `json:"side"`
	AmountB      string `json:"amount_b"`
	Symbol       string `json:"symbol"`
	ClientId     string `json:"client_id"`
}

func (p *Position) Transform() {
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.Side = strings.ToUpper(p.Side)
}

func (p *Position) Validate() error {
	if p.PositionSide != "LONG" && p.PositionSide != "SHORT" {
		return errors.New("positionSide must be LONG or SHORT")
	}

	if p.Side != "BUY" && p.Side != "SELL" {
		return errors.New("side must be BUY or SELL")
	}

	if p.AmountB == "" {
		return errors.New("amountB must be a number")
	}

	if p.Symbol == "" {
		return errors.New("symbol must be a string")
	}

	if p.ClientId == "" {
		return errors.New("clientId must be a string")
	}
	return nil
}

func (p *Position) ToDomain() domain.Position {
	processPosition := domain.Position{}
	processPosition.SetPositionSide(p.PositionSide)
	processPosition.SetSide(p.Side)
	processPosition.SetEntryQuantity(p.AmountB)
	processPosition.SetSymbol(p.Symbol)
	processPosition.SetClientId(p.ClientId)
	return processPosition
}
