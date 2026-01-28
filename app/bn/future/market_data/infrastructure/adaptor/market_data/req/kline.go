package req

type KlineRequest struct {
	Symbol    string `json:"symbol" binding:"required"`
	Interval  string `json:"interval" binding:"required"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
}
