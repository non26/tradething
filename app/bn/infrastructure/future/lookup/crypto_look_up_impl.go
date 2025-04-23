package infrastructure

import (
	"context"
	domainservice "tradething/app/bn/process/future/domain_service"

	future "tradething/app/bn/infrastructure/future"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type cryptoLookUp struct {
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository
}

func NewCryptoLookUp(bnFtCryptoTable bndynamodb.IBnFtCryptoRepository) future.ICryptoLookUp {
	return &cryptoLookUp{bnFtCryptoTable}
}

func (c *cryptoLookUp) LookUpBySymbol(ctx context.Context, symbol string, positionSide string) (*domainservice.LookUpSymbol, error) {

	cryptoCoin, err := c.bnFtCryptoTable.Get(ctx, symbol)
	if err != nil {
		return nil, err
	}

	if !cryptoCoin.IsFound() {
		cryptoCoin = dynamodbmodel.NewBinanceFutureCryptoTableRecord(symbol, positionSide)
	} else {
		cryptoCoin.SetNextCountingBy(positionSide)
	}

	lookup := &domainservice.LookUpSymbol{}
	lookup.SetSymbol(cryptoCoin.Symbol)
	lookup.SetCountingLong(cryptoCoin.CountingLong)
	lookup.SetCountingShort(cryptoCoin.CountingShort)
	return lookup, nil
}
