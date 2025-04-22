package domainservice

import (
	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type LookUpSymbol struct {
	symbol        string
	countingLong  int64
	countingShort int64
}

func NewLookUpSymbol() *LookUpSymbol {
	return &LookUpSymbol{}
}

func (l *LookUpSymbol) GetSymbol() string {
	return l.symbol
}

func (l *LookUpSymbol) GetCountingLong() int64 {
	return l.countingLong
}

func (l *LookUpSymbol) GetCountingShort() int64 {
	return l.countingShort
}

func (l *LookUpSymbol) SetSymbol(symbol string) {
	l.symbol = symbol
}

func (l *LookUpSymbol) SetCountingLong(countingLong int64) {
	l.countingLong = countingLong
}

func (l *LookUpSymbol) SetCountingShort(countingShort int64) {
	l.countingShort = countingShort
}

func (l *LookUpSymbol) GetCountingBy(positionSide string) int64 {
	if positionSide == bnconstant.LONG {
		return l.countingLong
	}
	return l.countingShort
}
