package infrastructure

import (
	"context"
	"errors"
	position "tradething/app/bn/infrastructure/future/position"
	domainservice "tradething/app/bn/process/future/domain_service"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type lookUp struct {
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewLookUp(
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) ILookUp {
	return &lookUp{
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (l *lookUp) ToOpeningPositionTable(position *position.Position) *dynamodbmodel.BnFtOpeningPosition {
	return &dynamodbmodel.BnFtOpeningPosition{
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		ClientId:     position.ClientId,
		Side:         position.Side,
		OrderType:    position.OrderType,
		AmountB:      position.AmountB,
	}
}

func (l *lookUp) LookUp(ctx context.Context, position *position.Position) (*domainservice.LookUp, error) {
	bnHistory, err := l.bnFtHistoryTable.Get(ctx, position.GetClientId())
	if err != nil {
		return nil, err
	}
	if bnHistory.IsFound() {
		return nil, errors.New("duplicate client id")
	}

	cryptoCoin, err := l.bnFtCryptoTable.Get(ctx, position.Symbol)
	if err != nil {
		return nil, err
	}

	if !cryptoCoin.IsFound() {
		cryptoCoin = dynamodbmodel.NewBinanceFutureCryptoTableRecord(position.Symbol, position.PositionSide)
	} else {
		cryptoCoin.SetNextCountingBy(position.PositionSide)
	}
	position.SetDefaultClientId(cryptoCoin.GetCountingBy(position.PositionSide))

	openingPosition, err := l.bnFtOpeningPositionTable.Get(ctx, l.ToOpeningPositionTable(position))
	if err != nil {
		return nil, err
	}

	lookUp := domainservice.NewLookUp()
	lookUp.OpeningPosition.SetIsFound(openingPosition.IsFound())
	lookUp.OpeningPosition.SetAmountB(openingPosition.AmountB)
	lookUp.OpeningPosition.SetClientId(openingPosition.ClientId)
	lookUp.Symbol.SetSymbol(cryptoCoin.Symbol)
	lookUp.Symbol.SetCountingLong(cryptoCoin.CountingLong)
	lookUp.Symbol.SetCountingShort(cryptoCoin.CountingShort)

	return lookUp, nil
}
