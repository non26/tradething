package model

import (
	model "tradething/app/kc/app/model/kcservicemodel/spot/order"
	"tradething/app/kc/kccommon"
)

type PlaceSpotOrderHandlerRequest struct {
	Side   string `json:"side"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Size   string `json:"size"`
	Funds  string `json:"funds"`
}

func (
	p *PlaceSpotOrderHandlerRequest,
) generateClientId() string {
	kccommon.CreateFutureClientId(
		"SPOT",
		p.Side,
		p.Symbol,
		"",
	)
	return ""
}

func (
	p *PlaceSpotOrderHandlerRequest,
) ToPlaceSpotOrderKcServiceRequest() *model.PlaceSpotOrderKcServiceRequest {
	m := &model.PlaceSpotOrderKcServiceRequest{
		ClientOid: kccommon.CreateFutureClientId("SPOT", p.Side, p.Symbol, ""),
		Side:      p.Side,
		Symbol:    p.Symbol,
		Type:      p.Type,
		Size:      p.Size,
		Funds:     p.Funds,
	}
	return m
}

type PlaceSpotOrderHandlerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
