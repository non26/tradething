package bnfuture

type CloseBySymbols struct {
	Data []CloseBySymbolsData
}

type CloseBySymbolsData struct {
	Symbol  string `json:"symbol"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
