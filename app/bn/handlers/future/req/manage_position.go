package req

type ManagePositionReq struct {
	CloseClientId            []string `json:"close_client_id"`
	AdvancedPositionClientId []string `json:"advanced_position_client_id"`
}
