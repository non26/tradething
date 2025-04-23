package process

import (
	"context"
	"errors"
	response "tradething/app/bn/handlers/future/res"
	domain "tradething/app/bn/process/future/domain"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

func (f *future) PlaceOrder(ctx context.Context, position domain.Position) (*response.Position, error) {

	bnposition := position.ToInfraPosition()
	tradeLookUp, err := f.infraTradeLookUp.LookUp(ctx, bnposition)
	if err != nil {
		return nil, err
	}

	cryptoLookUp, err := f.infraCryptoLookUp.LookUpBySymbol(ctx, position.GetSymbol(), position.GetPositionSide())
	if err != nil {
		return nil, err
	}
	bnposition.SetDefaultClientId(cryptoLookUp.GetCountingBy(position.GetPositionSide()))

	if utils.IsBuyPosition(position.GetSide(), position.GetPositionSide()) {
		if position.GetClientId() == tradeLookUp.OpeningPosition.GetClientId() {
			return nil, errors.New("duplicate client id")
		}
	}

	err = f.infraTrade.PlacePosition(ctx, bnposition)
	if err != nil {
		return nil, err
	}

	err = f.infraSavePosition.Save(ctx, bnposition, tradeLookUp)
	if err != nil {
		return nil, err
	}

	return &response.Position{
		ClientId: position.GetClientId(),
		Symbol:   position.GetSymbol(),
	}, nil
}
