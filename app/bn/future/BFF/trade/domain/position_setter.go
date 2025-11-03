package domain

import "strings"

func (p *Position) SetPositionSide(positionSide string) {
	p.positionSide = strings.ToUpper(positionSide)
}

func (p *Position) SetSide(side string) {
	p.side = strings.ToUpper(side)
}

func (p *Position) SetSymbol(symbol string) {
	p.symbol = strings.ToUpper(symbol)
}

func (p *Position) SetClientId(clientId string) {
	p.clientId = clientId
}

func (p *Position) SetMaxAccumulatePosition(maxAccumulatePosition string) {
	p.maxAccumulatePosition = maxAccumulatePosition
}

func (p *Position) SetAmountB(amountB string) {
	p.amountB = amountB
}
