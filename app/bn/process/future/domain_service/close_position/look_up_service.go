package closeposition

import (
	lookupitem "tradething/app/bn/process/future/domain_service"
)

type ClsoePositionLookUp struct {
	OpeningPosition *lookupitem.LookupOpeningPosition
}

func NewClsoePositionLookUp() *ClsoePositionLookUp {
	return &ClsoePositionLookUp{
		OpeningPosition: lookupitem.NewLookupOpeningPosition(),
	}
}
