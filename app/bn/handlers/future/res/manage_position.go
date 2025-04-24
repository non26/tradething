package res

type ManagePositionRes struct {
	ClosePosition    *CloseByClientIds `json:"close_position"`
	AdvancedPosition []*Position       `json:"advanced_position"`
}
