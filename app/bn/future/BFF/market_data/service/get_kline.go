package service

import (
	"context"
	"encoding/json"
	"fmt"
	"tradething/app/bn/future/BFF/market_data/domain"
)

func (s *service) GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error) {
	stringifyReq, _ := json.Marshal(req)
	fmt.Print("req from service", string(stringifyReq))
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
