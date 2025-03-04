package domainservice

import (
	lookupitem "tradething/app/bn/process/future/domain_service"
)

type TradeLookUp struct {
	OpeningPosition *lookupitem.LookupOpeningPosition
	Symbol          *lookupitem.LookUpSymbol
}

func NewTradeLookUp() *TradeLookUp {
	return &TradeLookUp{
		OpeningPosition: lookupitem.NewLookupOpeningPosition(),
		Symbol:          lookupitem.NewLookUpSymbol(),
	}
}
