package bnfuture

type InvalidatePosition struct {
	Result []InvalidatePositionData
}

type InvalidatePositionData struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
