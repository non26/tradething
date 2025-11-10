package res

type KlinesResponse struct {
	KlineData []KlineResponse
}

type KlineResponse struct {
	Open                  string `json:"open"`
	High                  string `json:"high"`
	Low                   string `json:"low"`
	Close                 string `json:"close"`
	IsCloseHigherThanOpen bool   `json:"is_green_candle"`
}
