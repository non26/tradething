package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type BnFtAccumulationRepository struct {
	db *dynamodb.Client
}

func NewBnFtAccumulationRepository(db *dynamodb.Client) *BnFtAccumulationRepository {
	return &BnFtAccumulationRepository{db: db}
}
