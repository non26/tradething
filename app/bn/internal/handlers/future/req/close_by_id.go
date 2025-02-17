package req

type ClosePositionByClientIds struct {
	ClientIds []string `json:"client_ids"`
}
