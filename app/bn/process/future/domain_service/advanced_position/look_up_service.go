package advancedposition

import domainservice "tradething/app/bn/process/future/domain_service"

type AdvancedPositionLookUp struct {
	AdvancedPosition *domainservice.LookUpAdvancedPosition
}

func NewAdvancedPositionLookUp() *AdvancedPositionLookUp {
	return &AdvancedPositionLookUp{
		AdvancedPosition: domainservice.NewLookUpAdvancedPosition(),
	}
}
