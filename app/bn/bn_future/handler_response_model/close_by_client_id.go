package bnfuture

type CloseByClientIds struct {
	Data []CloseByClientIdsData `json:"data"`
}

type CloseByClientIdsData struct {
	ClientId string `json:"client_id"`
	Status   string `json:"status"` // success or fail
	Message  string `json:"message"`
}
