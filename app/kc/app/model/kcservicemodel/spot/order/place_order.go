package model

type PlaceSpotOrderKcServiceRequest struct {
	ClientOid string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Funds     string `json:"funds"`
}

type PlaceSpotOrderKcServiceResponse struct {
	// Code    string `json:"code"`
	OrderId string `json:"orderId"`
}
