package externalservice

import (
	infraexternalservice "tradething/app/bn/future/BFF/market_data/infrastructure/external_service"
	externalservice "tradething/app/bn/future/market_data/service"
)

type marketData struct {
	service externalservice.IService
}

func NewMarketData(service externalservice.IService) infraexternalservice.IMarketData {
	return &marketData{service: service}
}
