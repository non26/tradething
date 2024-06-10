package model

type PlaceSignleOrderHandlerRequest struct {
	PositionSide     string  `json:"positionSide"`
	Side             string  `json:"side"`
	EntryQuantity    float64 `json:"entryQuantity"`
	Symbol           string  `json:"symbol"`
	LeverageLevel    int     `json:"leverageLevel"`
	NewClientOrderId string  `json:"newClientOrderId"`
	Timestamp        string  `json:"timestamp"`
}

type PlaceSignleOrderHandlerResponse struct {
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
}
