package bncommon

import "strings"

type IPositionSide interface {
	Long() string
	Short() string
	IsLong(side string) bool
	IsShort(side string) bool
}

type positionSide struct {
	long  string
	short string
}

func (p *positionSide) Long() string {
	return p.long
}

func (p *positionSide) Short() string {
	return p.short
}

func (p *positionSide) IsLong(side string) bool {
	return p.long == p.positionSideTransform(side)
}

func (p *positionSide) IsShort(side string) bool {
	return p.short == p.positionSideTransform(side)
}

func (p *positionSide) positionSideTransform(side string) string {
	return strings.ToUpper(side)
}

func NewPositionSide() IPositionSide {
	return &positionSide{
		long:  "LONG",
		short: "SHORT",
	}
}
