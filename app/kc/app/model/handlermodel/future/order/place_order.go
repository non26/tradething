package model

import (
	model "tradetoolv2/app/kc/app/model/kcservicemodel/future/order"
	"tradetoolv2/app/kc/kccommon"
)

type PlaceFutureOrderHandlerRequest struct {
	Side     string `json:"side"`
	Symbol   string `json:"symbol"`
	Leverage string `json:"leverage"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}

func (p *PlaceFutureOrderHandlerRequest) generateClientId() string {
	clietId := kccommon.CreateFutureClientId(
		"FUTURE",
		p.Side,
		p.Symbol,
		"",
	)
	return clietId
}

func (p *PlaceFutureOrderHandlerRequest) ToPlaceOrderKcServiceRequest() *model.PlaceFutureOrderKcServiceRequest {
	d := &model.PlaceFutureOrderKcServiceRequest{
		ClientOid: p.generateClientId(),
		Symbol:    p.Symbol,
		Leverage:  p.Leverage,
		Type:      p.Type,
		Side:      p.Side,
		Size:      p.Size,
	}
	return d
}

type PlaceFutureOrderHandlerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
