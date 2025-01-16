package bnfuture

import (
	"time"
	bntradereq "tradething/app/bn/bn_future/bnservice_request/trade"
)

type PositionSide struct {
	data []PositionSideInfo
}

func NewCloseBySymbolsServiceRequest() *PositionSide {
	return &PositionSide{
		data: []PositionSideInfo{},
	}
}

func (c *PositionSide) SetData(data PositionSideInfo) {
	c.data = append(c.data, data)
}

func (c *PositionSide) GetData() []PositionSideInfo {
	return c.data
}

type PositionSideInfo struct {
	symbol       string
	positionSide string
	amountQ      string
}

func (c *PositionSideInfo) SetSymbol(symbol string) {
	c.symbol = symbol
}

func (c *PositionSideInfo) GetSymbol() string {
	return c.symbol
}

func (c *PositionSideInfo) SetPositionSide(positionSide string) {
	c.positionSide = positionSide
}

func (c *PositionSideInfo) GetPositionSide() string {
	return c.positionSide
}

func (c *PositionSideInfo) SetAmountQ(amountQ string) {
	c.amountQ = amountQ
}

func (c *PositionSideInfo) GetAmountQ() string {
	return c.amountQ
}

func (p *PositionSideInfo) ToBinanceServiceModel(side string) *bntradereq.PlacePosition {
	m := bntradereq.PlacePosition{
		PositionSide:  p.positionSide,
		Side:          side,
		EntryQuantity: p.amountQ,
		Symbol:        p.symbol,
		ClientOrderId: "manual_close" + time.Now().In(time.UTC).Format(time.DateTime),
	}
	return &m
}
