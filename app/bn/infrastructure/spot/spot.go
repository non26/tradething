package infrastructure

import (
	"context"
	"tradething/app/bn/infrastructure/spot/order"
	domainservicecloseposition "tradething/app/bn/process/spot/domain_service/close_position"
	domainservicetrade "tradething/app/bn/process/spot/domain_service/trade"
)

type ITrade interface {
	PlaceOrder(ctx context.Context, order *order.Order) error
}

type ITradeLookUp interface {
	LookUp(ctx context.Context, order *order.Order) (*domainservicetrade.TradeLookUp, error)
}

type ITradeSaveOrder interface {
	Save(ctx context.Context, order *order.Order, lookup *domainservicetrade.TradeLookUp) error
}

type ICloseOrderLookUp interface {
	ById(ctx context.Context, clientId string) (*domainservicecloseposition.ClosePositionLookUp, error)
}
