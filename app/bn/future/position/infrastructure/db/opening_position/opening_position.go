package db

import (
	"tradething/app/bn/future/position/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type openingPositionRepository struct {
	db *dynamodb.Client
}

func NewOpeningPositionRepository(db *dynamodb.Client) db.IBnFtOpeningPositionRepository {
	return &openingPositionRepository{db: db}
}
