package domainservice

import (
	lookupitem "tradething/app/bn/process/spot/domain_service"
)

type TradeLookUp struct {
	OpeningPosition *lookupitem.LookUpOpeningPosition
	Symbol          *lookupitem.LookUpSymbol
}

func NewTradeLookUp() *TradeLookUp {
	return &TradeLookUp{
		OpeningPosition: lookupitem.NewLookUpOpeningPosition(),
		Symbol:          lookupitem.NewLookUpSymbol(),
	}
}
