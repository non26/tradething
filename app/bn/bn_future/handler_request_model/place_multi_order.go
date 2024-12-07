package bnfuture

type PlaceMultiOrderHandlerRequest struct {
	Orders []PlaceSignleOrderHandlerRequest `json:"orders"`
}

func (p *PlaceMultiOrderHandlerRequest) Transform() {
	for _, order := range p.Orders {
		order.Transform()
	}
}
