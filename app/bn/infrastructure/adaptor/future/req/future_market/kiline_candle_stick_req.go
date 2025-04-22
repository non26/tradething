package req

type KlineCandleStick struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}
