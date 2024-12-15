package bnfuture

import "strconv"

type price string

func (p price) GetFloat64() float64 {
	f, _ := strconv.ParseFloat(string(p), 64)
	return f
}

func (p price) Get() string {
	return string(p)
}

type CandleStickResponse [][]interface{}

func (c CandleStickResponse) GetOpenPrice() price {
	return price(c[0][1].(string))
}

func (c CandleStickResponse) GetClosePrice() price {
	return price(c[0][4].(string))
}
