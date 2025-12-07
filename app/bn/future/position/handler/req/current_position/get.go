package req

type GetCurrentPositionReq struct {
	ClientId     string `json:"client_id" validate:"required"`
	Symbol       string `json:"symbol" validate:"required"`
	AccountId    string `json:"account_id" validate:"required"`
	PositionSide string `json:"position_side" validate:"required"`
}
