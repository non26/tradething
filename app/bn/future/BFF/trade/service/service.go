package service

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor"
	accumapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/accumulation"
	positionapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position"
	positionhistoryapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position_history"
	subaccountapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/sub_account"
)

type ITradeService interface {
	NewOrder(ctx context.Context, order *domain.Order) error
	CloseOrderById(ctx context.Context, clientId string) error
	AccumulateOrder(ctx context.Context, order *domain.Order) error
}

type tradeService struct {
	tradeAdaptor            adaptor.ITradeAdaptor
	accumService            accumapi.IAccumulationExternalService
	currentPositionService  positionapi.ICurrentPositionExternalService
	advancedPositionService positionapi.IAdvancedPositionExternalService
	positionHistoryService  positionhistoryapi.IPositionHistoryExternalService
	subaccountService       subaccountapi.ISubAccountExternalService
}

func NewTradeService(
	tradeAdaptor adaptor.ITradeAdaptor,
	accumService accumapi.IAccumulationExternalService,
	currentPositionService positionapi.ICurrentPositionExternalService,
	advancedPositionService positionapi.IAdvancedPositionExternalService,
	positionHistoryService positionhistoryapi.IPositionHistoryExternalService,
	subaccountService subaccountapi.ISubAccountExternalService) ITradeService {
	return &tradeService{
		tradeAdaptor:            tradeAdaptor,
		accumService:            accumService,
		currentPositionService:  currentPositionService,
		advancedPositionService: advancedPositionService,
		positionHistoryService:  positionHistoryService,
		subaccountService:       subaccountService,
	}
}
