package service

import (
	"context"
	"tradething/app/bn/future/market_data/domain"
	"tradething/app/bn/future/market_data/handler/res"
)

func (s *service) GetPreviousKline(ctx context.Context, req *domain.Kline) (*res.KlinesResponse, error) {
	startTime, endTime := req.GetPreviousStartAndEndTimeInUnixTimestamp()
	req.StartTime = startTime
	req.EndTime = endTime
	kline, err := s.marketDataAdaptor.GetKline(ctx, req)
	if err != nil {
		return nil, err
	}
	data := kline.KlineData[0]
	klineResponse := &res.KlinesResponse{
		KlineData: []res.KlineResponse{
			{
				Open:                  data.Open,
				High:                  data.High,
				Low:                   data.Low,
				Close:                 data.Close,
				IsCloseHigherThanOpen: data.Close > data.Open,
			},
		},
	}
	return klineResponse, nil
}
