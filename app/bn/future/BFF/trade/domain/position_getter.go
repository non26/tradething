package domain

import "strings"

// getter
func (p *Position) GetPositionSide() string {
	return strings.ToUpper(p.positionSide)
}

func (p *Position) GetSide() string {
	return strings.ToUpper(p.side)
}

func (p *Position) GetAmountB() string {
	return p.amountB
}

func (p *Position) GetSymbol() string {
	return strings.ToUpper(p.symbol)
}

func (p *Position) GetClientId() string {
	return p.clientId
}

func (p *Position) GetMaxAccumulatePosition() string {
	return p.maxAccumulatePosition
}
