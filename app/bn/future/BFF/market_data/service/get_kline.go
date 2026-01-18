package service

import (
	"context"
	"tradething/app/bn/future/BFF/market_data/domain"
)

func (s *service) GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error) {
	kline, err := s.infra.GetKline(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(kline) == 0 {
		return nil, nil
	}
	if kline == nil {
		return nil, nil
	}
	return kline, nil
}
