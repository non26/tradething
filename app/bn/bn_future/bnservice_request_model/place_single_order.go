package bnfuture

import (
	"fmt"
	"strconv"
	"strings"
	"tradething/app/bn/bncommon"
)

type PlaceSignleOrderBinanceServiceRequest struct {
	PositionSide  string `json:"positionSide"`
	Side          string `json:"side"`
	EntryQuantity string `json:"quantity"`
	Symbol        string `json:"symbol"`
	ClientOrderId string `json:"newClientOrderId"`
	Type          string `json:"type"`
	Timestamp     string `json:"timestamp"`
}

func (p *PlaceSignleOrderBinanceServiceRequest) New() *PlaceSignleOrderBinanceServiceRequest {
	return &PlaceSignleOrderBinanceServiceRequest{}
}

func (p *PlaceSignleOrderBinanceServiceRequest) PrepareRequest() {
	p.Symbol = strings.ToUpper(p.Symbol)
	p.Side = strings.ToUpper(p.Side)
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.checkClientOrderId()
	p.checkOrderType()
	p.setTimestamp()
}

func (p *PlaceSignleOrderBinanceServiceRequest) GetData() interface{} {
	return p
}

func (p *PlaceSignleOrderBinanceServiceRequest) setTimestamp() {
	p.Timestamp = strconv.FormatInt(bncommon.GetTimeStamp(), 10)
}

func (p *PlaceSignleOrderBinanceServiceRequest) checkClientOrderId() {
	if p.ClientOrderId == "" {
		p.ClientOrderId = p.Symbol
	}
}

func (p *PlaceSignleOrderBinanceServiceRequest) checkOrderType() {
	if p.Type == "" {
		p.Type = "MARKET"
	}
}

func (p *PlaceSignleOrderBinanceServiceRequest) IsOrderTypeMarket() bool {
	return p.Type == "MARKET"
}

func (p *PlaceSignleOrderBinanceServiceRequest) IsOrderTypeLimit() bool {
	return p.Type == "LIMIT"
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetPositionSide(position string) {
	p.PositionSide = strings.ToUpper(position)
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetSide(side string) {
	p.Side = strings.ToUpper(side)
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetEntryQuantity(quantity string) {
	p.EntryQuantity = quantity
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetSymbol(symbol string) {
	p.Symbol = strings.ToUpper(symbol)
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetClientOrderId(client_order_id string) {
	p.ClientOrderId = client_order_id
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetDefaultClientOrderId(client_order_id string) {
	p.ClientOrderId = fmt.Sprintf("%v_%v", p.Symbol, client_order_id)
}

func (p *PlaceSignleOrderBinanceServiceRequest) SetType(order_type string) {
	p.Type = strings.ToUpper(order_type)
}

func NewPlaceSignleOrderBinanceServiceRequest(
	p *PlaceSignleOrderBinanceServiceRequest,
) IBnFutureServiceRequest {
	return p
}
