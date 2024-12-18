package bnfuture

type CloseBySymbolsHandlerResponse struct {
	Data []CloseBySymbolsHandlerResponseData `json:"data"`
}

type CloseBySymbolsHandlerResponseData struct {
	Symbol  string `json:"symbol"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
