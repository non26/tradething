package service

import (
	"context"
	"tradething/app/bn/future/market_data/domain"
)

func (s *service) GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error) {
	kline, err := s.marketDataAdaptor.GetKline(ctx, req)
	if err != nil {
		return nil, err
	}

	klineResponses := make([]*domain.Kline, len(kline.KlineData))
	for i, data := range kline.KlineData {
		klineResponses[i] = &domain.Kline{
			Open:          data.Open,
			High:          data.High,
			Low:           data.Low,
			Close:         data.Close,
			IsGreenCandle: data.Close > data.Open,
		}
	}
	return klineResponses, nil
}
