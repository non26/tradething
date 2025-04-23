package infrastructure

import (
	"context"
	position "tradething/app/bn/infrastructure/future/position"

	domainAdvPositionSvc "tradething/app/bn/process/future/domain_service/advanced_position"
	domainCryptoSvc "tradething/app/bn/process/future/domain_service/crypto"
	domainTradeSvc "tradething/app/bn/process/future/domain_service/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type saveBuyPosition struct {
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository
	bnFtCryptoTable          bndynamodb.IBnFtCryptoRepository
	bnFtHistoryTable         bndynamodb.IBnFtHistoryRepository
	bnFtAdvPositionTable     bndynamodb.IBnFtAdvancedPositionRepository
}

func NewSaveBuyPosition(
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvPositionTable bndynamodb.IBnFtAdvancedPositionRepository,
) ISavePositionBySide {
	return &saveBuyPosition{bnFtOpeningPositionTable, bnFtCryptoTable, bnFtHistoryTable, bnFtAdvPositionTable}
}

func (s *saveBuyPosition) Save(ctx context.Context, _position *position.Position, tradeLookUp *domainTradeSvc.TradeLookUp, cryptoLookUp *domainCryptoSvc.CryptoLookUp, advPositionLookUp *domainAdvPositionSvc.AdvancedPositionLookUp) error {

	// in case of accumulation
	if tradeLookUp.OpeningPosition.IsFound() {
		err := _position.AddMoreAmountB(tradeLookUp.OpeningPosition.GetAmountB())
		if err != nil {
			return err
		}
		err = s.bnFtHistoryTable.Insert(ctx, _position.ToHistoryTable())
		if err != nil {
			return err
		}
		_position.ClientId = tradeLookUp.OpeningPosition.GetClientId()
	}

	// in case of advacned position
	if advPositionLookUp.AdvancedPosition.IsFound() {
		err := s.bnFtAdvPositionTable.Delete(ctx, advPositionLookUp.AdvancedPosition.GetClientId())
		if err != nil {
			return err
		}
		advPosition := position.Position{
			ClientId:     advPositionLookUp.AdvancedPosition.GetClientId(),
			Symbol:       advPositionLookUp.AdvancedPosition.GetSymbol(),
			PositionSide: advPositionLookUp.AdvancedPosition.GetPositionSide(),
		}
		err = s.bnFtHistoryTable.Insert(ctx, advPosition.ToHistoryTable())
		if err != nil {
			return err
		}
	}

	err := s.bnFtOpeningPositionTable.Upsert(ctx, ToOpeningPositionTable(_position))
	if err != nil {
		return err
	}

	err = s.bnFtCryptoTable.Upsert(ctx, ToCryptoTable(cryptoLookUp))
	if err != nil {
		return err
	}

	return nil
}
