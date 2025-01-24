package bnfuture

import "github.com/shopspring/decimal"

type AccumulatePosition struct {
	accumulateId        string
	symbol              string
	positionSide        string
	side                string
	upperBoundPrice     float64
	lowerBoundPrice     float64
	totalAmountQ        string
	accumulationAmountQ string
}

func (m *AccumulatePosition) GetAccumulateId() string {
	return m.accumulateId
}

func (m *AccumulatePosition) SetAccumulateId(accumulateId string) {
	m.accumulateId = accumulateId
}

func (m *AccumulatePosition) GetSymbol() string {
	return m.symbol
}

func (m *AccumulatePosition) SetSymbol(symbol string) {
	m.symbol = symbol
}

func (m *AccumulatePosition) GetPositionSide() string {
	return m.positionSide
}

func (m *AccumulatePosition) SetPositionSide(positionSide string) {
	m.positionSide = positionSide
}

func (m *AccumulatePosition) GetSide() string {
	return m.side
}

func (m *AccumulatePosition) SetSide(side string) {
	m.side = side
}

func (m *AccumulatePosition) GetUpperBoundPrice() float64 {
	return m.upperBoundPrice
}

func (m *AccumulatePosition) SetUpperBoundPrice(upperBoundPrice float64) {
	m.upperBoundPrice = upperBoundPrice
}

func (m *AccumulatePosition) GetLowerBoundPrice() float64 {
	return m.lowerBoundPrice
}

func (m *AccumulatePosition) SetLowerBoundPrice(lowerBoundPrice float64) {
	m.lowerBoundPrice = lowerBoundPrice
}

func (m *AccumulatePosition) GetTotalAmountQ() string {
	return m.totalAmountQ
}

func (m *AccumulatePosition) SetTotalAmountQ(totalAmountQ string) {
	m.totalAmountQ = totalAmountQ
}

func (m *AccumulatePosition) GetAccumulationAmountQ() string {
	return m.accumulationAmountQ
}

func (m *AccumulatePosition) SetAccumulationAmountQ(accumulationAmountQ string) {
	m.accumulationAmountQ = accumulationAmountQ
}

func (m *AccumulatePosition) AddAccumulation(amount string) {
	total, _ := decimal.NewFromString(m.accumulationAmountQ)
	add, _ := decimal.NewFromString(amount)
	total = total.Add(add)
	m.accumulationAmountQ = total.String()
}
