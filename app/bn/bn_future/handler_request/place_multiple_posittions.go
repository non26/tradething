package bnfuture

import model "tradething/app/bn/bn_future/service_model"

type PlaceMultiplePositions struct {
	Positions []PlacePosition `json:"positions"`
}

func (p *PlaceMultiplePositions) Transform() {
	for _, order := range p.Positions {
		order.Transform()
	}
}

func (p *PlaceMultiplePositions) ToServiceModel() *model.Positions {
	orders := make([]model.Position, 0, len(p.Positions))
	for _, order := range p.Positions {
		orders = append(orders, *order.ToServiceModel())
	}
	return &model.Positions{
		Positions: orders,
	}
}
