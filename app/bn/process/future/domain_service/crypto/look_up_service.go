package domainservice

import (
	lookupItem "tradething/app/bn/process/future/domain_service"
)

type CryptoLookUp struct {
	Symbol *lookupItem.LookUpSymbol
}

func NewCryptoLookUp() *CryptoLookUp {
	return &CryptoLookUp{
		Symbol: lookupItem.NewLookUpSymbol(),
	}
}
