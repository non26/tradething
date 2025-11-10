package service

import (
	"tradething/app/bn/future/position/infrastructure/db"
	"tradething/app/bn/future/position/service"
)

type advancedPositionService struct {
	repository db.IBnFtAdvancedPositionRepository
}

func NewAdvancedPositionService(repository db.IBnFtAdvancedPositionRepository) service.IAdvancedPositionService {
	return &advancedPositionService{repository: repository}
}
