package externalservice

import (
	"context"
	"tradething/app/bn/future/BFF/market_data/domain"
	"tradething/app/bn/future/BFF/market_data/infrastructure/external_service/market_data/dto"
)

func (m *marketData) GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error) {
	reqdto := dto.NewEmptyKline()
	reqdto = reqdto.FromDomain(req)
	kline, err := m.service.GetKline(ctx, reqdto.ToExternalServiceDomain())
	if err != nil {
		return nil, err
	}

	if len(kline) == 0 {
		return nil, nil
	}

	if kline == nil {
		return nil, nil
	}

	klineResponses := make([]*domain.Kline, len(kline))
	for i, data := range kline {
		klineresponse := dto.NewEmptyKline()
		klineresponse = klineresponse.FromExternalServiceDomain(data)
		klineResponses[i] = klineresponse.ToDomain()
	}
	return klineResponses, nil
}
