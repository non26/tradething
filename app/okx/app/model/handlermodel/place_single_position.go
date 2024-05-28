package model

import okxservicemodel "tradetoolv2/app/okx/app/model/okxservicemodel"

type PlaceASinglePositionHandlerRequest struct {
	InstId  string `json:"instId"`
	TdMode  string `json:"tdMode"`
	Side    string `json:"side"`
	PosSide string `json:"posSide"`
	OrdType string `json:"ordType"`
	Sz      string `json:"sz"`
	TgtCcy  string `json:"tgtCcy"`
}

func (
	p *PlaceASinglePositionHandlerRequest,
) ToPlaceASinglePositionOKXServiceRequest() *okxservicemodel.PlaceASinglePositionOKXServiceRequest {
	okxmodel := &okxservicemodel.PlaceASinglePositionOKXServiceRequest{
		InstId:  p.InstId,
		TdMode:  p.TdMode,
		Side:    p.Side,
		PosSide: p.PosSide,
		OrdType: p.OrdType,
		Sz:      p.Sz,
		TgtCcy:  p.TgtCcy,
	}
	return okxmodel
}

type PlaceASinglePositionHandlerResponse struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

func (
	p *PlaceASinglePositionHandlerResponse,
) ToPlaceASinglePositionHandlerRequest(okxresponse *okxservicemodel.PlaceASinglePositionOKXserviceResponse) {
	p.ClOrdId = okxresponse.ClOrdId
	p.OrdId = okxresponse.OrdId
	p.Tag = okxresponse.Tag
	p.SCode = okxresponse.SCode
	p.SMsg = okxresponse.SMsg
}
