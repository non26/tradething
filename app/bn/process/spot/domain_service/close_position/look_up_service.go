package domainservice

import (
	lookupitem "tradething/app/bn/process/spot/domain_service"
	domainservice "tradething/app/bn/process/spot/domain_service/trade"
)

type ClosePositionLookUp struct {
	OpeningPosition *lookupitem.LookUpOpeningPosition
}

func NewClosePositionLookUp() *ClosePositionLookUp {
	return &ClosePositionLookUp{
		OpeningPosition: lookupitem.NewLookUpOpeningPosition(),
	}
}

func (c *ClosePositionLookUp) ToTradeLookUp() *domainservice.TradeLookUp {
	return &domainservice.TradeLookUp{
		OpeningPosition: c.OpeningPosition,
	}
}
