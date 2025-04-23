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
	tradeLookUp, err := f.infraTradeLookUp.LookUp(ctx, bnposition)
	if err != nil {
		return nil, err
	}
	if tradeLookUp.OpeningPosition.IsFound() {
		return nil, errors.New("duplicate history client id")
	}

	advancedPositionLookUp, err := f.infraAdvancedPositionLookUp.LookUpByClientId(ctx, position.GetClientId())
	if err != nil {
		return nil, err
	}
	if advancedPositionLookUp.AdvancedPosition.IsFound() {
		position = domain.NewPosition(
			position.GetClientId(),
			position.GetSymbol(),
			position.GetPositionSide(),
			position.GetSide(),
			position.GetEntryQuantity(),
		)
	}

	cryptoLookUp, err := f.infraCryptoLookUp.LookUpBySymbol(ctx, position.GetSymbol(), position.GetPositionSide())
	if err != nil {
		return nil, err
	}
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

	err = f.infraSavePosition.Save(ctx, bnposition, tradeLookUp, cryptoLookUp, advancedPositionLookUp)
	if err != nil {
		return nil, err
	}

	return &response.Position{
		ClientId: position.GetClientId(),
		Symbol:   position.GetSymbol(),
	}, nil
}
