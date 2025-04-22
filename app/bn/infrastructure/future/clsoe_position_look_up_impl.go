package infrastructure

// import (
// 	"context"
// 	"errors"
// 	domainservice "tradething/app/bn/process/future/domain_service/close_position"

// 	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
// )

// type closePositionLookUp struct {
// 	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
// 	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
// 	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
// }

// func NewClosePositionLookUp(
// 	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
// 	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
// 	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
// ) IClosePositionLookup {
// 	return &closePositionLookUp{bnFtOpeningPositionTable, bnFtCryptoTable, bnFtHistoryTable}
// }

// func (c *closePositionLookUp) ById(ctx context.Context, clientId string) (*domainservice.ClsoePositionLookUp, error) {
// 	bnHistory, err := c.bnFtHistoryTable.Get(ctx, clientId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if bnHistory.IsFound() {
// 		return nil, errors.New("position already closed")
// 	}

// 	bnOpening, err := c.bnFtOpeningPositionTable.ScanWith(ctx, clientId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !bnOpening.IsFound() {
// 		return nil, errors.New("position not found")
// 	}

// 	lookUp := domainservice.NewClsoePositionLookUp()
// 	lookUp.OpeningPosition.SetIsFound(bnOpening.IsFound())
// 	lookUp.OpeningPosition.SetAmountB(bnOpening.AmountB)
// 	lookUp.OpeningPosition.SetClientId(bnOpening.ClientId)
// 	lookUp.OpeningPosition.SetSymbol(bnOpening.Symbol)
// 	lookUp.OpeningPosition.SetPositionSide(bnOpening.PositionSide)
// 	return lookUp, nil
// }
