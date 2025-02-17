package domain

import (
	"strings"

	infra "tradething/app/bn/internal/infrastructure/future/position"

	positionconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type Position struct {
	positionSide string
	side         string
	amountB      string
	symbol       string
	clientId     string
}

func (p *Position) GetPositionSide() string {
	return p.positionSide
}

func (p *Position) SetPositionSide(positionSide string) {
	p.positionSide = strings.ToUpper(positionSide)
}

func (p *Position) GetSide() string {
	return p.side
}

func (p *Position) SetSide(side string) {
	p.side = strings.ToUpper(side)
}

func (p *Position) GetEntryQuantity() string {
	return p.amountB
}

func (p *Position) SetEntryQuantity(entryQuantity string) {
	p.amountB = entryQuantity
}

func (p *Position) GetSymbol() string {
	return p.symbol
}

func (p *Position) SetSymbol(symbol string) {
	p.symbol = symbol
}

func (p *Position) GetClientId() string {
	return p.clientId
}

func (p *Position) SetClientId(clientId string) {
	p.clientId = clientId
}

func (p *Position) IsLongPosition() bool {
	return p.positionSide == positionconstant.LONG
}

func (p *Position) IsShortPosition() bool {
	return p.positionSide == positionconstant.SHORT
}

func (p *Position) IsBuyOrder() bool {
	if p.IsLongPosition() && p.side == positionconstant.BUY {
		return true
	}
	if p.IsShortPosition() && p.side == positionconstant.SELL {
		return true
	}
	return false
}

func (p *Position) IsSellOrder() bool {
	if p.IsLongPosition() && p.side == positionconstant.SELL {
		return true
	}
	if p.IsShortPosition() && p.side == positionconstant.BUY {
		return true
	}
	return false
}

func (p *Position) ToInfraPosition() *infra.Position {
	infraPosition := infra.Position{
		PositionSide: p.positionSide,
		AmountB:      p.amountB,
		Symbol:       p.symbol,
		// OrderType    : p.Order
		ClientId: p.clientId,
		Side:     p.side,
	}
	return &infraPosition
}
