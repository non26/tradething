package req

import "tradething/app/bn/future/market_data/domain"

type GetKlineReq struct {
	Symbol    string `json:"symbol" binding:"required"`
	Interval  string `json:"interval" binding:"required"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

func (r *GetKlineReq) ToDomain() *domain.Kline {
	return &domain.Kline{
		Symbol:    r.Symbol,
		Interval:  r.Interval,
		StartTime: r.StartTime,
		EndTime:   r.EndTime,
	}
}
