package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
	domain "tradething/app/bn/process/future/domain"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

func (f *future) PlaceOrder(ctx context.Context, position *domain.Position) (*response.Position, error) {

	bnposition := position.ToInfraPosition()
	// look up opening position and history position
	tradeLookUp, err := f.infraTradeLookUp.LookUp(ctx, bnposition)
	if err != nil {
		return nil, err
	}
	if tradeLookUp.OpeningPosition.IsFound() {
		if tradeLookUp.OpeningPosition.GetClientId() == position.GetClientId() {
			return nil, errors.New("duplicate opening position")
		}
		// return nil, errors.New("duplicate opening position")
	}

	// look up advanced position
	advancedPositionLookUp, err := f.infraAdvancedPositionLookUp.LookUpByClientId(ctx, position.GetClientId())
	if err != nil {
		return nil, err
	}
	if advancedPositionLookUp.AdvancedPosition.IsFound() {
		position = domain.NewPositionWith(
			advancedPositionLookUp.AdvancedPosition.GetClientId(),
			advancedPositionLookUp.AdvancedPosition.GetSymbol(),
			advancedPositionLookUp.AdvancedPosition.GetPositionSide(),
			advancedPositionLookUp.AdvancedPosition.GetSide(),
			advancedPositionLookUp.AdvancedPosition.GetAmountB(),
		)
	}

	// look up crypto coin
	cryptoLookUp, err := f.infraCryptoLookUp.LookUpBySymbol(ctx, position.GetSymbol(), position.GetPositionSide())
	if err != nil {
		return nil, err
	}
	bnposition = position.ToInfraPosition()
	bnposition.SetDefaultClientId(cryptoLookUp.Symbol.GetCountingBy(position.GetPositionSide()))

	if utils.IsBuyPosition(position.GetSide(), position.GetPositionSide()) {
		if position.GetClientId() == tradeLookUp.OpeningPosition.GetClientId() {
			return nil, errors.New("duplicate client id")
		}
	}

	err = f.infraTrade.PlacePosition(ctx, bnposition)
	if err != nil {
		return nil, err
	}

	// after success opeing position in Binance, save position to database
	err = f.infraSavePosition.Save(ctx, bnposition, tradeLookUp, cryptoLookUp, advancedPositionLookUp)
	if err != nil {
		return nil, err
	}

	return &response.Position{
		ClientId: position.GetClientId(),
		Symbol:   position.GetSymbol(),
	}, nil
}
