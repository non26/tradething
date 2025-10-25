package domain

import (
	"strings"

	infra "tradething/app/bn/infrastructure/future/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/shopspring/decimal"
)

type Position struct {
	positionSide          string
	side                  string
	amountB               string
	symbol                string
	clientId              string
	maxAccumulatePosition string
}

func NewPositionWith(clientId string, symbol string, positionSide string, side string, amountB string) *Position {
	return &Position{
		clientId:     clientId,
		symbol:       symbol,
		positionSide: positionSide,
		side:         side,
		amountB:      amountB,
	}
}

func NewPosition() *Position {
	return &Position{}
}

func (p *Position) GetPositionSide() string {
	return p.positionSide
}

func (p *Position) SetPositionSide(positionSide string) *Position {
	p.positionSide = strings.ToUpper(positionSide)
	return p
}

func (p *Position) GetSide() string {
	return p.side
}

func (p *Position) SetSide(side string) *Position {
	p.side = strings.ToUpper(side)
	return p
}

func (p *Position) GetEntryQuantity() string {
	return p.amountB
}

func (p *Position) SetEntryQuantity(entryQuantity string) *Position {
	p.amountB = entryQuantity
	return p
}

func (p *Position) GetSymbol() string {
	return p.symbol
}

func (p *Position) SetSymbol(symbol string) *Position {
	p.symbol = symbol
	return p
}

func (p *Position) GetClientId() string {
	return p.clientId
}

func (p *Position) SetClientId(clientId string) *Position {
	p.clientId = clientId
	return p
}

func (p *Position) GetMaxAccumulatePosition() string {
	return p.maxAccumulatePosition
}

func (p *Position) SetMaxAccumulatePosition(maxAccumulatePosition string) *Position {
	p.maxAccumulatePosition = maxAccumulatePosition
	return p
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

func (p *Position) ExceedMaxAccumulatePosition() bool {
	if p.maxAccumulatePosition != "" {
		_maxAccumulatePosition, err := decimal.NewFromString(p.maxAccumulatePosition)
		if err != nil {
			return false
		}
		_amountB, err := decimal.NewFromString(p.amountB)
		if err != nil {
			return false
		}
		return _amountB.GreaterThan(_maxAccumulatePosition)
	}
	return false
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

func (p *Position) ToInfraPositionWith(clientId string, symbol string, positionSide string, side string, amountB string) *infra.Position {
	infraPosition := infra.Position{
		PositionSide: positionSide,
		AmountB:      amountB,
		Symbol:       symbol,
		OrderType:    bnconstant.MARKET,
		ClientId:     clientId,
		Side:         side,
	}
	return &infraPosition
}

func (p *Position) NewClosePositionFrom(
	clientId string,
	symbol string,
	positionSide string,
	amountB string) *Position {
	newPosition := Position{}
	newPosition.SetClientId(clientId)
	newPosition.SetSymbol(symbol)
	newPosition.SetPositionSide(positionSide)
	newPosition.SetSellSideFrom(positionSide)
	newPosition.SetEntryQuantity(amountB)
	return &newPosition
}
