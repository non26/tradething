package model

import okxservicemodel "tradetoolv2/app/okx/app/model/okxservicemodel"

type SetNewLeverageHandlerRequest struct {
	InstId  string `json:"instId"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	PosSide string `json:"posSide"`
}

func (
	s *SetNewLeverageHandlerRequest,
) ToSetNewLeverageOKXServiceRequest() *okxservicemodel.SetNewLeverageOKXServiceRequest {
	okxmodel := &okxservicemodel.SetNewLeverageOKXServiceRequest{
		InstId:  s.InstId,
		Lever:   s.Lever,
		MgnMode: s.MgnMode,
	}
	return okxmodel
}

type SetNewLeverageHandlerResponse struct {
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}

func (
	s *SetNewLeverageHandlerResponse,
) ToSetNewLeverageHandlerResponse(
	okxresponse *okxservicemodel.SetNewLeverageOKXserviceResponse,
) {
	s.Lever = okxresponse.Lever
	s.MgnMode = okxresponse.MgnMode
	s.InstId = okxresponse.InstId
	s.PosSide = okxresponse.PosSide
}
