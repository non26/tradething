package infrastructure

import (
	"context"
	adaptor "tradething/app/bn/infrastructure/adaptor/future"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type ShortPosition struct {
	BnTradeservice           adaptor.IBinanceFutureTradeService
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewShortPosition(
	BnTradeservice adaptor.IBinanceFutureTradeService,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) IPosition {
	return &ShortPosition{
		BnTradeservice:           BnTradeservice,
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (p *ShortPosition) BuyPosition(ctx context.Context, position *Position) error {
	// cryptoCoin, err := p.bnFtCryptoTable.Get(ctx, position.Symbol)
	// if err != nil {
	// 	return err
	// }
	// if !cryptoCoin.IsFound() {
	// 	cryptoCoin = dynamodbmodel.NewBinanceFutureCryptoTableRecord(position.Symbol, position.PositionSide)
	// } else {
	// 	cryptoCoin.SetCountingShort(cryptoCoin.GetNextCountingShort().Int())
	// }
	// position.SetDefaultClientId(cryptoCoin.GetCountingShort())

	// openingPosition, err := p.bnFtOpeningPositionTable.Get(ctx, p.ToOpeningPositionTable(position))
	// if err != nil {
	// 	return err
	// }
	// if position.ClientId == openingPosition.ClientId {
	// 	return errors.New("duplicate client id")
	// }

	err := p.placePosition(ctx, position)
	if err != nil {
		return err
	}

	// if openingPosition.IsFound() {
	// 	err = position.AddMoreAmountB(openingPosition.AmountB)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = p.bnFtHistoryTable.Insert(ctx, position.ToHistoryTable())
	// 	if err != nil {
	// 		return err
	// 	}
	// 	position.ClientId = openingPosition.ClientId
	// }
	// // this would be Upsert
	// err = p.bnFtOpeningPositionTable.Upsert(ctx, p.ToOpeningPositionTable(position))
	// if err != nil {
	// 	return err
	// }

	// err = p.bnFtCryptoTable.Upsert(ctx, cryptoCoin)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (p *ShortPosition) SellPosition(ctx context.Context, position *Position) error {
	err := p.placePosition(ctx, position)
	if err != nil {
		return err
	}

	// err = p.bnFtOpeningPositionTable.Delete(ctx, p.ToOpeningPositionTable(position))
	// if err != nil {
	// 	return err
	// }

	// err = p.bnFtHistoryTable.Insert(ctx, p.ToHistoryTable(position))
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (p *ShortPosition) placePosition(ctx context.Context, position *Position) error {
	_, err := p.BnTradeservice.PlaceOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}
	return nil
}

func (p *ShortPosition) ToOpeningPositionTable(position *Position) *dynamodbmodel.BnFtOpeningPosition {
	return &dynamodbmodel.BnFtOpeningPosition{
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		ClientId:     position.ClientId,
		Side:         position.Side,
		OrderType:    position.OrderType,
		AmountB:      position.AmountB,
	}
}

func (p *ShortPosition) ToHistoryTable(position *Position) *dynamodbmodel.BnFtHistory {
	return &dynamodbmodel.BnFtHistory{
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		ClientId:     position.ClientId,
	}
}
