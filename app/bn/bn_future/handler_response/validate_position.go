package bnfuture

type ValidatePosition struct {
	Data []ValidatePositionData
}

type ValidatePositionData struct {
	ClientId string `json:"client_id"`
	Status   string `json:"status"` // success or fail
	Message  string `json:"message"`
}
