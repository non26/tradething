package domain

import (
	"strings"

	infra "tradething/app/bn/infrastructure/future/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
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
	return p.positionSide == bnconstant.LONG
}

func (p *Position) IsShortPosition() bool {
	return p.positionSide == bnconstant.SHORT
}

func (p *Position) IsBuyOrder() bool {
	if p.IsLongPosition() && p.side == bnconstant.BUY {
		return true
	}
	if p.IsShortPosition() && p.side == bnconstant.SELL {
		return true
	}
	return false
}

func (p *Position) IsSellOrder() bool {
	if p.IsLongPosition() && p.side == bnconstant.SELL {
		return true
	}
	if p.IsShortPosition() && p.side == bnconstant.BUY {
		return true
	}
	return false
}

func (p *Position) SetSellSideFrom(positionSide string) {
	if positionSide == bnconstant.LONG {
		p.side = bnconstant.SELL
	} else {
		p.side = bnconstant.BUY
	}
}

func (p *Position) ToInfraPosition() *infra.Position {
	infraPosition := infra.Position{
		PositionSide: p.positionSide,
		AmountB:      p.amountB,
		Symbol:       p.symbol,
		OrderType:    bnconstant.MARKET,
		ClientId:     p.clientId,
		Side:         p.side,
	}
	return &infraPosition
}
