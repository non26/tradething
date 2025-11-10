package service

import (
	"context"
	"tradething/app/bn/future/market_data/domain"
)

func (s *service) GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error) {
	startTime, endTime := req.GetPreviousStartAndEndTimeInUnixTimestamp()
	req.StartTime = startTime
	req.EndTime = endTime
	kline, err := s.marketDataAdaptor.GetKline(ctx, req)
	if err != nil {
		return nil, err
	}
	data := kline.KlineData[0]
	klineResponse := &domain.Kline{
		Open:          data.Open,
		High:          data.High,
		Low:           data.Low,
		Close:         data.Close,
		IsGreenCandle: data.Close > data.Open,
	}
	return klineResponse, nil
}
