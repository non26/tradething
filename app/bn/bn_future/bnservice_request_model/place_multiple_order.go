package bnfuture

import (
	"strconv"
	"strings"

	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

type PlaceMultiOrderBinanceServiceRequest struct {
	BatchOrders []BatchOrder `json:"batchOrders"`
	Timestamp   string       `json:"timestamp"`
}

type BatchOrder struct {
	PositionSide  string `json:"positionSide"`
	Side          string `json:"side"`
	EntryQuantity string `json:"quantity"`
	Symbol        string `json:"symbol"`
	ClientOrderId string `json:"newClientOrderId"`
	Type          string `json:"type"`
}

func (p *PlaceMultiOrderBinanceServiceRequest) setTimestamp() {
	p.Timestamp = strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
}

func (p *BatchOrder) checkClientOrderId() {
	if p.ClientOrderId == "" {
		p.ClientOrderId = p.Symbol
	}
}

func (p *BatchOrder) checkOrderType() {
	if p.Type == "" {
		p.Type = "MARKET"
	}
}

func (p *PlaceMultiOrderBinanceServiceRequest) PrepareRequest() {
	for _, order := range p.BatchOrders {
		order.Symbol = strings.ToUpper(order.Symbol)
		order.Side = strings.ToUpper(order.Side)
		order.PositionSide = strings.ToUpper(order.PositionSide)
		order.checkClientOrderId()
		order.checkOrderType()
	}
	p.setTimestamp()
}

func (p *PlaceMultiOrderBinanceServiceRequest) GetData() interface{} {
	return p
}

func NewPlaceMultiOrderBinanceServiceRequest(
	request *PlaceMultiOrderBinanceServiceRequest,
) IBnFutureServiceRequest {
	return request
}
