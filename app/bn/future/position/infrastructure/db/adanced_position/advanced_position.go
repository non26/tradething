package db

import (
	"tradething/app/bn/future/position/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type advancedPositionRepository struct {
	db *dynamodb.Client
}

func NewAdvancedPositionRepository(db *dynamodb.Client) db.IBnFtAdvancedPositionRepository {
	return &advancedPositionRepository{db: db}
}
