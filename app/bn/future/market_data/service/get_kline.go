package service

import (
	"context"
	"tradething/app/bn/future/market_data/domain"
	"tradething/app/bn/future/market_data/handler/res"
)

func (s *service) GetKline(ctx context.Context, req *domain.Kline) (*res.KlinesResponse, error) {
	kline, err := s.marketDataAdaptor.GetKline(ctx, req)
	if err != nil {
		return nil, err
	}

	klineResponse := &res.KlinesResponse{
		KlineData: []res.KlineResponse{},
	}
	for _, data := range kline.KlineData {
		klineResponse.KlineData = append(klineResponse.KlineData, res.KlineResponse{
			Open:                  data.Open,
			High:                  data.High,
			Low:                   data.Low,
			Close:                 data.Close,
			IsCloseHigherThanOpen: data.Close > data.Open,
		})
	}
	return klineResponse, nil
}
