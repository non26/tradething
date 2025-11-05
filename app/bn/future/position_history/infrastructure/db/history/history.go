package db

import (
	"tradething/app/bn/future/position_history/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type bnFtHistoryRepository struct {
	db *dynamodb.Client
}

func NewBnFtHistoryRepository(db *dynamodb.Client) db.IBnFtHistoryRepository {
	return &bnFtHistoryRepository{db: db}
}
