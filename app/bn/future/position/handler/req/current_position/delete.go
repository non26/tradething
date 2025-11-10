package req

type DeleteCurrentPositionReq struct {
	ClientId     string `json:"client_id" binding:"required"`
	Symbol       string `json:"symbol" binding:"required"`
	AccountId    string `json:"account_id" binding:"required"`
	PositionSide string `json:"position_side" binding:"required"`
}
