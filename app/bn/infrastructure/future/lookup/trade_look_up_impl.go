package infrastructure

import (
	"context"
	"errors"
	position "tradething/app/bn/infrastructure/future/position"
	domainservice "tradething/app/bn/process/future/domain_service/trade"

	future "tradething/app/bn/infrastructure/future"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type tradeLookUp struct {
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
}

func NewTradeLookUp(
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) future.ITradeLookUp {
	return &tradeLookUp{
		bnFtOpeningPositionTable: bnFtOpeningPositionTable,
		bnFtCryptoTable:          bnFtCryptoTable,
		bnFtHistoryTable:         bnFtHistoryTable,
	}
}

func (l *tradeLookUp) ToOpeningPositionTable(position *position.Position) *dynamodbmodel.BnFtOpeningPosition {
	return &dynamodbmodel.BnFtOpeningPosition{
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		ClientId:     position.ClientId,
		Side:         position.Side,
		OrderType:    position.OrderType,
		AmountB:      position.AmountB,
	}
}

func (l *tradeLookUp) LookUp(ctx context.Context, position *position.Position) (*domainservice.TradeLookUp, error) {
	bnHistory, err := l.bnFtHistoryTable.Get(ctx, position.GetClientId())
	if err != nil {
		return nil, err
	}
	if bnHistory.IsFound() {
		return nil, errors.New("duplicate history client id")
	}

	var openingPosition *dynamodbmodel.BnFtOpeningPosition
	if position.Symbol != "" && position.PositionSide != "" {
		openingPosition, err = l.bnFtOpeningPositionTable.Get(ctx, l.ToOpeningPositionTable(position))
		if err != nil {
			return nil, err
		}
	} else {
		openingPosition, err = l.bnFtOpeningPositionTable.ScanWith(ctx, position.ClientId)
		if err != nil {
			return nil, err
		}
	}

	lookUp := domainservice.NewTradeLookUp()
	lookUp.OpeningPosition.SetIsFound(openingPosition.IsFound())
	lookUp.OpeningPosition.SetAmountB(openingPosition.AmountB)
	lookUp.OpeningPosition.SetClientId(openingPosition.ClientId)
	lookUp.OpeningPosition.SetSymbol(openingPosition.Symbol)
	lookUp.OpeningPosition.SetPositionSide(openingPosition.PositionSide)
	lookUp.OpeningPosition.SetSide(openingPosition.Side)

	return lookUp, nil
}
