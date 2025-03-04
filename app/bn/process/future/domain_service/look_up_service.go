package domainservice

type LookUp struct {
	OpeningPosition *lookupOpeningPosition
	Symbol          *lookUpSymbol
}

func NewLookUp() *LookUp {
	return &LookUp{
		OpeningPosition: NewLookupOpeningPosition(),
		Symbol:          NewLookUpSymbol(),
	}
}
