package req

type DeleteAdvancedPositionReq struct {
	ClientId string `json:"client_id" binding:"required"`
}
