package spot

import (
	"context"
	"errors"
	infrastructure "tradething/app/internal/infrastructure/spot"
	"tradething/app/internal/process/spot/domain"

	res "tradething/app/internal/handlers/spot/res"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type ISpot interface {
	PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error)
	CloseOrderByClientIds(ctx context.Context, clientIds []string) (*res.CloseByClientIds, error)
	MultiplePosition(ctx context.Context, positions []domain.Order) (*res.MultipleOrder, error)
}

type spot struct {
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
	infraSpot                  infrastructure.ITrade
}

func NewSpot(
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
	infraSpot infrastructure.ITrade,
) ISpot {
	return &spot{
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
		infraSpot:                  infraSpot,
	}
}

func (s *spot) PlaceOrder(ctx context.Context, order *domain.Order) (*res.Trade, error) {

	spotHistory, err := s.bnSpotHistoryTable.Get(ctx, order.ClientId)
	if err != nil {
		return nil, err
	}

	if spotHistory.IsFound() {
		return nil, errors.New("duplicate client id")
	}

	err = s.infraSpot.PlaceOrder(ctx, order.ToInfrastructureOrder())
	if err != nil {
		return nil, err
	}

	return &res.Trade{
		ClientId: order.ClientId,
		Symbol:   order.Symbol,
	}, nil
}

func (s *spot) CloseOrderByClientIds(ctx context.Context, clientIds []string) (response *res.CloseByClientIds, err error) {
	for _, clientId := range clientIds {
		spotHistory, err := s.bnSpotHistoryTable.Get(ctx, clientId)
		if err != nil {
			return nil, err
		}
		if spotHistory.IsFound() {
			response.AddWithData(clientId, "", "client id isfound")
			continue
		}

		openingPosition, err := s.bnSpotOpeningPositionTable.ScanWith(ctx, clientId)
		if err != nil {
			return nil, err
		}

		if openingPosition.IsFound() {
			response.AddWithData(clientId, "", "opening position isfound")
			continue
		}
		orderRequest := domain.Order{
			ClientId: clientId,
			Symbol:   openingPosition.Symbol,
			Side:     bnconstant.SELL,
			AmountB:  openingPosition.AmountB,
		}
		err = s.infraSpot.PlaceOrder(ctx, orderRequest.ToInfrastructureOrder())
		if err != nil {
			response.AddWithData(clientId, orderRequest.Symbol, "failed")
			continue
		}
		response.AddWithData(clientId, orderRequest.Symbol, "success")
	}
	return response, nil
}

func (s *spot) MultiplePosition(ctx context.Context, orders []domain.Order) (response *res.MultipleOrder, err error) {
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
