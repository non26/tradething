package infrastructure

import (
	"context"

	domainCryptoSvc "tradething/app/bn/process/future/domain_service/crypto"

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

func (c *cryptoLookUp) LookUpBySymbol(ctx context.Context, symbol string, positionSide string) (*domainCryptoSvc.CryptoLookUp, error) {

	cryptoCoin, err := c.bnFtCryptoTable.Get(ctx, symbol)
	if err != nil {
		return nil, err
	}

	if !cryptoCoin.IsFound() {
		cryptoCoin = dynamodbmodel.NewBinanceFutureCryptoTableRecord(symbol, positionSide)
	} else {
		cryptoCoin.SetNextCountingBy(positionSide)
	}

	lookup := &domainCryptoSvc.CryptoLookUp{}
	lookup.Symbol.SetSymbol(cryptoCoin.Symbol)
	lookup.Symbol.SetCountingLong(cryptoCoin.CountingLong)
	lookup.Symbol.SetCountingShort(cryptoCoin.CountingShort)
	return lookup, nil
}
