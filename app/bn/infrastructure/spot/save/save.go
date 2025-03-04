package save

import (
	"context"
	"tradething/app/bn/infrastructure/spot/order"
	domainservice "tradething/app/bn/process/spot/domain_service/trade"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_spot/models"
)

type ISaveOrder interface {
	Save(ctx context.Context, order *order.Order, lookup *domainservice.TradeLookUp) error
}

func ToCrypto(symbol string, counting int64) *dynamodbmodel.BnSpotCrypto {
	return &dynamodbmodel.BnSpotCrypto{
		Symbol:   symbol,
		Counting: counting,
	}
}
