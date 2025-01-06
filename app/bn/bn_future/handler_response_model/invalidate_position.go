package bnfuture

type InvalidatePositionHandlerResponse struct {
	Result []InvalidatePositionHandlerResponseData `json:"result"`
}

type InvalidatePositionHandlerResponseData struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
