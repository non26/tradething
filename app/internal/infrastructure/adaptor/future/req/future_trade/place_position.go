package req

import (
	"fmt"
	"strconv"
	"strings"

	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnutils "github.com/non26/tradepkg/pkg/bn/utils"
)

type PlacePosition struct {
	PositionSide  string `json:"positionSide"`
	Side          string `json:"side"`
	EntryQuantity string `json:"quantity"`
	Symbol        string `json:"symbol"`
	ClientOrderId string `json:"newClientOrderId"`
	Type          string `json:"type"`
	Timestamp     string `json:"timestamp"`
}

func (p *PlacePosition) New() *PlacePosition {
	return &PlacePosition{}
}

func (p *PlacePosition) PrepareRequest() {
	p.Symbol = strings.ToUpper(p.Symbol)
	p.Side = strings.ToUpper(p.Side)
	p.PositionSide = strings.ToUpper(p.PositionSide)
	p.checkClientOrderId()
	p.checkOrderType()
	p.setTimestamp()
}

func (p *PlacePosition) GetData() interface{} {
	return p
}

func (p *PlacePosition) setTimestamp() {
	p.Timestamp = strconv.FormatInt(bnutils.GetBinanceTimestamp(), 10)
}

func (p *PlacePosition) checkClientOrderId() {
	if p.ClientOrderId == "" {
		p.ClientOrderId = p.Symbol
	}
}

func (p *PlacePosition) checkOrderType() {
	if p.Type == "" {
		p.Type = "MARKET"
	}
}

func (p *PlacePosition) IsOrderTypeMarket() bool {
	return p.Type == "MARKET"
}

func (p *PlacePosition) IsOrderTypeLimit() bool {
	return p.Type == "LIMIT"
}

func (p *PlacePosition) SetPositionSide(position string) {
	p.PositionSide = strings.ToUpper(position)
}

func (p *PlacePosition) SetSide(side string) {
	p.Side = strings.ToUpper(side)
}

func (p *PlacePosition) SetEntryQuantity(quantity string) {
	p.EntryQuantity = quantity
}

func (p *PlacePosition) SetSymbol(symbol string) {
	p.Symbol = strings.ToUpper(symbol)
}

func (p *PlacePosition) SetClientOrderId(client_order_id string) {
	p.ClientOrderId = client_order_id
}

func (p *PlacePosition) SetDefaultClientOrderId(client_order_id string) {
	p.ClientOrderId = fmt.Sprintf("%v_%v", p.Symbol, client_order_id)
}

func (p *PlacePosition) SetType(order_type string) {
	p.Type = strings.ToUpper(order_type)
}

func NewPlaceSignleOrderBinanceServiceRequest(
	p *PlacePosition,
) bnrequest.IBnFutureServiceRequest {
	return p
}
