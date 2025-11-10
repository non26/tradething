package req

type KlineRequest struct {
	Symbol    string `json:"symbol" binding:"required"`
	Interval  string `json:"interval" binding:"required"`
	StartTime int64  `json:"startTime" binding:"required"`
	EndTime   int64  `json:"endTime" binding:"required"`
}
