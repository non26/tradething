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
	// cryptoCoin, err := l.bnFtCryptoTable.Get(ctx, position.Symbol)
	// if err != nil {
	// 	return err
	// }
	// if !cryptoCoin.IsFound() {
	// 	cryptoCoin = dynamodbmodel.NewBinanceFutureCryptoTableRecord(position.Symbol, position.PositionSide)
	// } else {
	// 	cryptoCoin.SetCountingLong(cryptoCoin.GetNextCountingLong().Int())
	// }
	// position.SetDefaultClientId(cryptoCoin.GetCountingLong())

	// openingPosition, err := l.bnFtOpeningPositionTable.Get(ctx, l.ToOpeningPositionTable(position))
	// if err != nil {
	// 	return err
	// }
	// if position.ClientId == openingPosition.ClientId {
	// 	return errors.New("duplicate client id")
	// }

	err := l.placePosition(ctx, position)
	if err != nil {
		return err
	}

	// if openingPosition.IsFound() {
	// 	err = position.AddMoreAmountB(openingPosition.AmountB)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	err = l.bnFtHistoryTable.Insert(ctx, position.ToHistoryTable())
	// 	if err != nil {
	// 		return err
	// 	}
	// 	position.ClientId = openingPosition.ClientId
	// }
	// // this would be Upsert
	// err = l.bnFtOpeningPositionTable.Upsert(ctx, l.ToOpeningPositionTable(position))
	// if err != nil {
	// 	return err
	// }
	// // upsert
	// err = l.bnFtCryptoTable.Upsert(ctx, cryptoCoin)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (l *LongPosition) SellPosition(ctx context.Context, position *Position) error {
	err := l.placePosition(ctx, position)
	if err != nil {
		return err
	}

	// err = l.bnFtOpeningPositionTable.Delete(ctx, l.ToOpeningPositionTable(position))
	// if err != nil {
	// 	return err
	// }

	// err = l.bnFtHistoryTable.Insert(ctx, l.ToHistoryTable(position))
	// if err != nil {
	// 	return err
	// }

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
