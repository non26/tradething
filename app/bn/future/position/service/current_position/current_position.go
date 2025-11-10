package service

import (
	"tradething/app/bn/future/position/infrastructure/db"
	"tradething/app/bn/future/position/service"
)

type currentPositionService struct {
	repository db.IBnFtOpeningPositionRepository
}

func NewCurrentPositionService(repository db.IBnFtOpeningPositionRepository) service.ICurrentPositionService {
	return &currentPositionService{repository: repository}
}
