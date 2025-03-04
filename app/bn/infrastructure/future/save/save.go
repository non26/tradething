package save

import (
	"context"

	position "tradething/app/bn/infrastructure/future/position"
	domainservice "tradething/app/bn/process/future/domain_service"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type ISavePositionBySide interface {
	Save(ctx context.Context, position *position.Position, lookup *domainservice.LookUp) error
}

func ToOpeningPositionTable(position *position.Position) *dynamodbmodel.BnFtOpeningPosition {
	return &dynamodbmodel.BnFtOpeningPosition{
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
		ClientId:     position.ClientId,
		Side:         position.Side,
		OrderType:    position.OrderType,
		AmountB:      position.AmountB,
	}
}

func ToHistoryTable(position *position.Position) *dynamodbmodel.BnFtHistory {
	return &dynamodbmodel.BnFtHistory{
		ClientId:     position.ClientId,
		Symbol:       position.Symbol,
		PositionSide: position.PositionSide,
	}
}

func ToCryptoTable(lookUp *domainservice.LookUp) *dynamodbmodel.BnFtCrypto {
	return &dynamodbmodel.BnFtCrypto{
		Symbol:        lookUp.Symbol.GetSymbol(),
		CountingLong:  lookUp.Symbol.GetCountingLong(),
		CountingShort: lookUp.Symbol.GetCountingShort(),
	}
}
