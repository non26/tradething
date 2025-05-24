package req

type ManagePositionReq struct {
	ClosePosition    []string `json:"close_position"`
	AdvancedPosition []string `json:"advanced_position"`
}
