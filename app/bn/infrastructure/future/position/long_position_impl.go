package infrastructure

import (
	"context"
	adaptor "tradething/app/bn/infrastructure/adaptor/future"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type LongPosition struct {
	bnTradeservice           adaptor.IBinanceFutureTradeService
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewLongPosition(
	bnTradeservice adaptor.IBinanceFutureTradeService,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) IPosition {
	return &LongPosition{
		bnTradeservice:           bnTradeservice,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (l *LongPosition) BuyPosition(ctx context.Context, position *Position) error {
	err := l.placePosition(ctx, position)
	if err != nil {
		return err
	}
	return nil
}

func (l *LongPosition) SellPosition(ctx context.Context, position *Position) error {
	err := l.placePosition(ctx, position)
	if err != nil {
		return err
	}
	return nil
}

func (p *LongPosition) placePosition(ctx context.Context, position *Position) error {
	_, err := p.bnTradeservice.PlaceOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}
	return nil
}

func (l *LongPosition) ToOpeningPositionTable(position *Position) *dynamodbmodel.BnFtOpeningPosition {
	return &dynamodbmodel.BnFtOpeningPosition{
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		ClientId:     position.ClientId,
		Side:         position.Side,
		OrderType:    position.OrderType,
		AmountB:      position.AmountB,
	}
}

func (l *LongPosition) ToHistoryTable(position *Position) *dynamodbmodel.BnFtHistory {
	return &dynamodbmodel.BnFtHistory{
		ClientId:     position.ClientId,
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
	}
}
