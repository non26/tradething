package bncommon

import "strings"

type IOrderType interface {
	Market() string
	Limit() string
	Stop() string
	StopLimit() string
	IsMarket(order_type string) bool
	IsLimit(order_type string) bool
	IsStop(order_type string) bool
	IsStopLimit(order_type string) bool
}

type orderType struct {
	market    string
	limit     string
	stop      string
	stopLimit string
}

func (o *orderType) Market() string {
	return o.market
}

func (o *orderType) Limit() string {
	return o.limit
}

func (o *orderType) Stop() string {
	return o.stop
}

func (o *orderType) StopLimit() string {
	return o.stopLimit
}

func (o *orderType) IsMarket(order_type string) bool {
	return o.market == o.orderTypeTransform(order_type)
}

func (o *orderType) IsLimit(order_type string) bool {
	return o.limit == o.orderTypeTransform(order_type)
}

func (o *orderType) IsStop(order_type string) bool {
	return o.stop == o.orderTypeTransform(order_type)
}

func (o *orderType) IsStopLimit(order_type string) bool {
	return o.stopLimit == o.orderTypeTransform(order_type)
}

func (o *orderType) orderTypeTransform(order_type string) string {
	return strings.ToUpper(order_type)
}

func NewOrderType() IOrderType {
	return &orderType{
		market:    "MARKET",
		limit:     "LIMIT",
		stop:      "STOP",
		stopLimit: "STOP_LIMIT",
	}
}
