package infrastructure

import (
	"context"
	"errors"
	adaptor "tradething/app/bn/internal/infrastructure/adaptor/future"

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
	qouteUsdt, err := p.bnFtCryptoTable.Get(ctx, position.Symbol)
	if err != nil {
		return err
	}
	if qouteUsdt == nil {
		qouteUsdt = dynamodbmodel.NewBinanceFutureCryptoTableRecord(position.Symbol, position.PositionSide)
		err := p.bnFtCryptoTable.Insert(ctx, qouteUsdt)
		if err != nil {
			return err
		}
	} else {
		qouteUsdt.SetCountingShort(qouteUsdt.GetNextCountingShort().Int())
	}
	position.SetDefaultClientId(qouteUsdt.GetCountingShort())

	openingPosition, err := p.bnFtOpeningPositionTable.Get(ctx, p.ToOpeningPositionTable(position))
	if err != nil {
		return err
	}
	if position.ClientId == openingPosition.ClientId {
		return errors.New("duplicate client id")
	}

	err = p.placePosition(ctx, position)
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
	err = p.bnFtOpeningPositionTable.Upsert(ctx, p.ToOpeningPositionTable(position))
	if err != nil {
		return err
	}

	err = p.bnFtCryptoTable.Update(ctx, qouteUsdt)
	if err != nil {
		return err
	}

	return nil
}

func (p *ShortPosition) SellPosition(ctx context.Context, position *Position) error {
	err := p.placePosition(ctx, position)
	if err != nil {
		return err
	}

	err = p.bnFtOpeningPositionTable.Delete(ctx, p.ToOpeningPositionTable(position))
	if err != nil {
		return err
	}

	err = p.bnFtHistoryTable.Insert(ctx, p.ToHistoryTable(position))
	if err != nil {
		return err
	}

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
	}
}
