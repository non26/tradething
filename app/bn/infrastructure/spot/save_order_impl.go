package infrastructure

import (
	"context"
	"tradething/app/bn/infrastructure/spot/order"
	"tradething/app/bn/infrastructure/spot/save"
	domainservice "tradething/app/bn/process/spot/domain_service/trade"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type ITradeSaveOrderBuilder interface {
	Get(ctx context.Context, order *order.Order) save.ISaveOrder
}

type tradeSaveOrderBuilder struct {
	saveBuy  save.ISaveOrder
	saveSell save.ISaveOrder
}

func NewTradeSaveOrderBuilder(
	saveBuy save.ISaveOrder,
	saveSell save.ISaveOrder,
) ITradeSaveOrderBuilder {
	return &tradeSaveOrderBuilder{
		saveBuy:  saveBuy,
		saveSell: saveSell,
	}
}

func (t *tradeSaveOrderBuilder) Get(ctx context.Context, order *order.Order) save.ISaveOrder {
	if order.Side == bnconstant.BUY {
		return t.saveBuy
	}
	return t.saveSell
}

type tradeSaveOrder struct {
	tradeSaveOrderBuilder      ITradeSaveOrderBuilder
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewTradeSaveOrder(
	tradeSaveOrderBuilder ITradeSaveOrderBuilder,
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ITradeSaveOrder {
	return &tradeSaveOrder{
		tradeSaveOrderBuilder:      tradeSaveOrderBuilder,
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (t *tradeSaveOrder) Save(ctx context.Context, order *order.Order, lookup *domainservice.TradeLookUp) error {
	saveOrder := t.tradeSaveOrderBuilder.Get(ctx, order)
	return saveOrder.Save(ctx, order, lookup)
}
