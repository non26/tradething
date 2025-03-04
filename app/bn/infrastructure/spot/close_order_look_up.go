package infrastructure

import (
	"context"
	"errors"

	domainservicecloseposition "tradething/app/bn/process/spot/domain_service/close_position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

type closeOrderLookUp struct {
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository
	bnSpotCryptoTable          bndynamodb.IBnSpotCryptoRepository
	bnSpotHistoryTable         bndynamodb.IBnSpotHistoryRepository
}

func NewCloseOrderLookUp(
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotCryptoTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
) ICloseOrderLookUp {
	return &closeOrderLookUp{
		bnSpotOpeningPositionTable: bnSpotOpeningPositionTable,
		bnSpotCryptoTable:          bnSpotCryptoTable,
		bnSpotHistoryTable:         bnSpotHistoryTable,
	}
}

func (c *closeOrderLookUp) ById(ctx context.Context, clientId string) (*domainservicecloseposition.ClosePositionLookUp, error) {
	spotHistory, err := c.bnSpotHistoryTable.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	if spotHistory.IsFound() {
		return nil, errors.New("client id is found")
	}

	openingPosition, err := c.bnSpotOpeningPositionTable.ScanWith(ctx, clientId)
	if err != nil {
		return nil, err
	}

	if !openingPosition.IsFound() {
		return nil, errors.New("opening position is not found")
	}

	lookup := domainservicecloseposition.NewClosePositionLookUp()
	lookup.OpeningPosition.SetIsFound(openingPosition.IsFound())
	lookup.OpeningPosition.SetClientId(openingPosition.ClientId)
	lookup.OpeningPosition.SetSymbol(openingPosition.Symbol)
	lookup.OpeningPosition.SetQuantity(openingPosition.AmountB)

	return lookup, nil
}
