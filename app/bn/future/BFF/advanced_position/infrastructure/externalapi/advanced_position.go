package externalapi

import (
	"tradething/app/bn/future/position/service"
)

type advancedPositionExternalService struct {
	advancedPositionService service.IAdvancedPositionService
}

type IAdvancedPositionExternalService interface {
	GetAdvancedPosition() service.IAdvancedPositionService
}

func NewAdvancedPositionExternalService(advancedPositionService service.IAdvancedPositionService) IAdvancedPositionExternalService {
	return &advancedPositionExternalService{advancedPositionService: advancedPositionService}
}

func (s *advancedPositionExternalService) GetAdvancedPosition() service.IAdvancedPositionService {
	return s.advancedPositionService
}
