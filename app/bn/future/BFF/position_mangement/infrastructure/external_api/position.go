package externalapi

import (
	"tradething/app/bn/future/position/service"
)

type positionExternalService struct {
	currentPositionService  service.ICurrentPositionService
	advancedPositionService service.IAdvancedPositionService
}

type IPositionExternalService interface {
	GetPosition() service.ICurrentPositionService
	GetAdvancedPosition() service.IAdvancedPositionService
}

func NewPositionService(currentPositionService service.ICurrentPositionService, advancedPositionService service.IAdvancedPositionService) IPositionExternalService {
	return &positionExternalService{currentPositionService: currentPositionService, advancedPositionService: advancedPositionService}
}

func (s *positionExternalService) GetPosition() service.ICurrentPositionService {
	return s.currentPositionService
}

func (s *positionExternalService) GetAdvancedPosition() service.IAdvancedPositionService {
	return s.advancedPositionService
}
