package domain

import (
	"tradething/app/bn/infrastructure/spot/order"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type Order struct {
	Symbol   string
	Side     string
	AmountB  string
	ClientId string
}

func (o *Order) ToInfrastructureOrder() *order.Order {
	return &order.Order{
		Symbol:           o.Symbol,
		Side:             o.Side,
		Quantity:         o.AmountB,
		NewClientOrderId: o.ClientId,
		Type:             bnconstant.MARKET,
	}
}
