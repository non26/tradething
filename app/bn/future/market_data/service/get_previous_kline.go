package service

import (
	"context"
	"errors"
	"strconv"
	"tradething/app/bn/future/market_data/domain"
)

func (s *service) GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error) {
	startTime, endTime := req.GetPreviousStartAndEndTimeInUnixTimestamp()
	req.StartTime = strconv.FormatInt(startTime, 10)
	req.EndTime = strconv.FormatInt(endTime, 10)
	kline, err := s.marketDataAdaptor.GetKline(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(kline.KlineData) == 0 {
		return nil, errors.New("no previous kline data")
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
