package closeposition

import (
	lookupitem "tradething/app/bn/process/future/domain_service"
	domainservice "tradething/app/bn/process/future/domain_service/trade"
)

type ClsoePositionLookUp struct {
	OpeningPosition *lookupitem.LookupOpeningPosition
}

func NewClsoePositionLookUp() *ClsoePositionLookUp {
	return &ClsoePositionLookUp{
		OpeningPosition: lookupitem.NewLookupOpeningPosition(),
	}
}

func (c *ClsoePositionLookUp) ToTradeLookUp() *domainservice.TradeLookUp {
	return &domainservice.TradeLookUp{
		OpeningPosition: c.OpeningPosition,
	}
}
