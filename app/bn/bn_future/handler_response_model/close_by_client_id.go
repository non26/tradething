package bnfuture

type CloseByClientIdsHandlerResponse struct {
	Data []CloseByClientIdsHandlerResponseData `json:"data"`
}

type CloseByClientIdsHandlerResponseData struct {
	ClientId string `json:"client_id"`
	Status   string `json:"status"` // success or fail
	Message  string `json:"message"`
}
