package externalservice

import (
	"context"
	"tradething/app/bn/future/BFF/market_data/domain"
	"tradething/app/bn/future/BFF/market_data/infrastructure/external_service/market_data/dto"
)

func (m *marketData) GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error) {
	reqdto := dto.NewEmptyKline()

	kline, err := m.service.GetPreviousKline(ctx, reqdto.ToExternalServiceDomain())
	if err != nil {
		return nil, err
	}

	if kline == nil {
		return nil, nil
	}

	klineresponse := dto.NewEmptyKline()
	klineresponse = klineresponse.FromExternalServiceDomain(kline)
	return klineresponse.ToDomain(), nil
}
