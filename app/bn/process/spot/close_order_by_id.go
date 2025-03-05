package spot

import (
	"context"
	res "tradething/app/bn/handlers/spot/res"
	"tradething/app/bn/process/spot/domain"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func (s *spot) CloseOrderByClientIds(ctx context.Context, clientIds []string) (response *res.CloseByClientIds, err error) {
	response = &res.CloseByClientIds{}
	for _, clientId := range clientIds {
		lookup, err := s.infraClosePositionLookUp.ById(ctx, clientId)
		if err != nil {
			response.AddWithData(clientId, "", err.Error())
			continue
		}

		orderRequest := domain.Order{
			ClientId: clientId,
			Symbol:   lookup.OpeningPosition.GetSymbol(),
			Side:     bnconstant.SELL,
			AmountB:  lookup.OpeningPosition.GetQuantity(),
		}
		err = s.infraSpot.PlaceOrder(ctx, orderRequest.ToInfrastructureOrder())
		if err != nil {
			response.AddWithData(clientId, orderRequest.Symbol, "failed")
			continue
		}

		err = s.infraSpotSaveOrder.Save(ctx, orderRequest.ToInfrastructureOrder(), lookup.ToTradeLookUp())
		if err != nil {
			response.AddWithData(clientId, orderRequest.Symbol, "failed")
			continue
		}

		response.AddWithData(clientId, orderRequest.Symbol, "success")
	}
	return response, nil
}
