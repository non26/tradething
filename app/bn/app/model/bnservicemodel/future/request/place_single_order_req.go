package bnserivcemodelreq

import (
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

func (p *PlaceSignleOrderBinanceServiceRequest) PrepareRequest() {
	p.Symbol = strings.ToUpper(p.Symbol)
	p.Side = strings.ToUpper(p.Side)
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.checkClientOrderId()
	p.checkOrderType()
	p.setTimestamp()
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

func NewPlaceSignleOrderBinanceServiceRequest(
	p *PlaceSignleOrderBinanceServiceRequest,
) IBnFutureServiceRequest {
	return p
}
