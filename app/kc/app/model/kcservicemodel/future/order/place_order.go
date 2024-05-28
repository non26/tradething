package model

type PlaceFutureOrderKcServiceRequest struct {
	ClientOid string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol"`
	Leverage  string `json:"leverage"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
}

type PlaceFutureOrderKcServiceData struct {
	OrderID string `json:"orderId"`
}

type PlaceFutureOrderKcServiceResponse struct {
	Code string                        `json:"code"`
	Data PlaceFutureOrderKcServiceData `json:"data"`
}
