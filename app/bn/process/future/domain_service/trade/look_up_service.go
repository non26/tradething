package domainservice

import (
	lookupItem "tradething/app/bn/process/future/domain_service"
)

type TradeLookUp struct {
	OpeningPosition *lookupItem.LookupOpeningPosition
}

func NewTradeLookUp() *TradeLookUp {
	return &TradeLookUp{
		OpeningPosition: lookupItem.NewLookupOpeningPosition(),
	}
}
