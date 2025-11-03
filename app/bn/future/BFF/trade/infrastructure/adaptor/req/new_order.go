package req

type NewOrderRequest struct {
	AccountId        string `json:"accountId" binding:"required"`
	PositionSide     string `json:"positionSide" binding:"required"`
	Side             string `json:"side" binding:"required"`
	Quantity         string `json:"quantity" binding:"required"`
	Symbol           string `json:"symbol" binding:"required"`
	NewClientOrderId string `json:"newClientOrderId" binding:"required"`
	Type             string `json:"type" binding:"required"`
}
