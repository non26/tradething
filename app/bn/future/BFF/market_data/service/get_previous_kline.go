package service

import (
	"context"
	"tradething/app/bn/future/BFF/market_data/domain"
)

func (s *service) GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error) {

	kline, err := s.infra.GetPreviousKline(ctx, req)
	if err != nil {
		return nil, err
	}
	if kline == nil {
		return nil, nil
	}
	return kline, nil
}
