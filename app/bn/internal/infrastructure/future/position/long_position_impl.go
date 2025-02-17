package infrastructure

import (
	"context"
	"errors"
	adaptor "tradething/app/bn/internal/infrastructure/adaptor/future"

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
	qouteUsdt, err := l.bnFtCryptoTable.Get(ctx, position.Symbol)
	if err != nil {
		return err
	}
	if !qouteUsdt.IsFound() {
		qouteUsdt = dynamodbmodel.NewBinanceFutureCryptoTableRecord(position.Symbol, position.PositionSide)
		err := l.bnFtCryptoTable.Insert(ctx, qouteUsdt)
		if err != nil {
			return err
		}
	} else {
		qouteUsdt.SetCountingLong(qouteUsdt.GetNextCountingLong().Int())
	}
	position.SetDefaultClientId(qouteUsdt.GetCountingLong())

	openingPosition, err := l.bnFtOpeningPositionTable.Get(ctx, l.ToOpeningPositionTable(position))
	if err != nil {
		return err
	}
	if position.ClientId == openingPosition.ClientId {
		return errors.New("duplicate client id")
	}

	err = l.placePosition(ctx, position)
	if err != nil {
		return err
	}

	if openingPosition.IsFound() {
		err = position.AddMoreAmountB(openingPosition.AmountB)
		if err != nil {
			return err
		}
	}
	// this would be Upsert
	err = l.bnFtOpeningPositionTable.Upsert(ctx, l.ToOpeningPositionTable(position))
	if err != nil {
		return err
	}

	err = l.bnFtCryptoTable.Update(ctx, qouteUsdt)
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

	err = l.bnFtOpeningPositionTable.Delete(ctx, l.ToOpeningPositionTable(position))
	if err != nil {
		return err
	}

	err = l.bnFtHistoryTable.Insert(ctx, l.ToHistoryTable(position))
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
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
	}
}
