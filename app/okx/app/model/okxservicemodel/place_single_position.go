package model

type PlaceASinglePositionOKXServiceRequest struct {
	InstId  string `json:"instId"`
	TdMode  string `json:"tdMode"`
	Side    string `json:"side"`
	PosSide string `json:"posSide"`
	OrdType string `json:"ordType"`
	Sz      string `json:"sz"`
	Px      string `json:"px"`
	PxUsd   string `json:"pxUsd"`
	PxVol   string `json:"pxVol"`
	TgtCcy  string `json:"tgtCcy"`
}

type PlaceASinglePositionOKXserviceResponse struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
