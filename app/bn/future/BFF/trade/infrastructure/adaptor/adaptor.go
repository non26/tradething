package adaptor

import (
	"context"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor/req"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor/res"
)

type IAdaptor interface {
	NewOrder(ctx context.Context, req *req.NewOrderRequest) (*res.NewOrderResponse, error)
}

type adaptor struct {
	baseUrl          string
	newOrderEndpoint string
}

func NewAdaptor(baseUrl string, newOrderEndpoint string) IAdaptor {
	return &adaptor{
		baseUrl:          baseUrl,
		newOrderEndpoint: newOrderEndpoint,
	}
}
