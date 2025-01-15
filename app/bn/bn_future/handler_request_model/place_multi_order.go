package bnfuture

import svcfuture "tradething/app/bn/bn_future/service_model"

type PlaceMultiOrderHandlerRequest struct {
	Orders []PlaceSignleOrderHandlerRequest `json:"orders"`
}

func (p *PlaceMultiOrderHandlerRequest) Transform() {
	for _, order := range p.Orders {
		order.Transform()
	}
}

func (p *PlaceMultiOrderHandlerRequest) ToServiceModel() *svcfuture.PlaceMultiOrderServiceRequest {
	orders := make([]svcfuture.Position, 0, len(p.Orders))
	for _, order := range p.Orders {
		orders = append(orders, *order.ToServiceModel())
	}
	return &svcfuture.PlaceMultiOrderServiceRequest{
		Positions: orders,
	}
}
