package spot

import (
	"context"
	infrastructure "tradething/app/bn/infrastructure/spot"
	"tradething/app/bn/process/spot/domain"

	res "tradething/app/bn/handlers/spot/res"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type ISpot interface {
	PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error)
	CloseOrderByClientIds(ctx context.Context, clientIds []string) (*res.CloseByClientIds, error)
	MultiplePosition(ctx context.Context, positions []domain.Order) (*res.MultipleOrder, error)
}

type spot struct {
	infraSpotLookUp            infrastructure.ITradeLookUp
	infraSpotSaveOrder         infrastructure.ITradeSaveOrder
	infraClosePositionLookUp   infrastructure.ICloseOrderLookUp
	infraSpot                  infrastructure.ITrade
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewSpot(
	infraSpotLookUp infrastructure.ITradeLookUp,
	infraSpotSaveOrder infrastructure.ITradeSaveOrder,
	infraClosePositionLookUp infrastructure.ICloseOrderLookUp,
	infraSpot infrastructure.ITrade,
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ISpot {
	return &spot{
		infraSpotLookUp:            infraSpotLookUp,
		infraSpotSaveOrder:         infraSpotSaveOrder,
		infraClosePositionLookUp:   infraClosePositionLookUp,
		infraSpot:                  infraSpot,
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (s *spot) PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error) {

	lookup, err := s.infraSpotLookUp.LookUp(ctx, order.ToInfrastructureOrder())
	if err != nil {
		return nil, err
	}

	err = s.infraSpot.PlaceOrder(ctx, order.ToInfrastructureOrder())
	if err != nil {
		return nil, err
	}

	err = s.infraSpotSaveOrder.Save(ctx, order.ToInfrastructureOrder(), lookup)
	if err != nil {
		return nil, err
	}

	return &res.Trade{
		ClientId: order.ClientId,
		Symbol:   order.Symbol,
	}, nil
}

func (s *spot) CloseOrderByClientIds(ctx context.Context, clientIds []string) (response *res.CloseByClientIds, err error) {
	response = &res.CloseByClientIds{}
	for _, clientId := range clientIds {
		lookup, err := s.infraClosePositionLookUp.ById(ctx, clientId)
		if err != nil {
			response.AddWithData(clientId, "", err.Error())
			continue
		}

		orderRequest := domain.Order{
			ClientId: clientId,
			Symbol:   lookup.OpeningPosition.GetSymbol(),
			Side:     bnconstant.SELL,
			AmountB:  lookup.OpeningPosition.GetQuantity(),
		}
		err = s.infraSpot.PlaceOrder(ctx, orderRequest.ToInfrastructureOrder())
		if err != nil {
			response.AddWithData(clientId, orderRequest.Symbol, "failed")
			continue
		}

		err = s.infraSpotSaveOrder.Save(ctx, orderRequest.ToInfrastructureOrder(), lookup.ToTradeLookUp())
		if err != nil {
			response.AddWithData(clientId, orderRequest.Symbol, "failed")
			continue
		}

		response.AddWithData(clientId, orderRequest.Symbol, "success")
	}
	return response, nil
}

func (s *spot) MultiplePosition(ctx context.Context, orders []domain.Order) (response *res.MultipleOrder, err error) {
	response = &res.MultipleOrder{}
	for _, order := range orders {
		_, err := s.PlaceOrder(ctx, &order)
		if err != nil {
			response.AddWithData(order.ClientId, order.Symbol, "failed")
			continue
		}
		response.AddWithData(order.ClientId, order.Symbol, "success")
	}
	return response, nil
}
