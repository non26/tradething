package db

import (
	itfdb "tradething/app/bn/future/sub_account/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type registerAccountRepository struct {
	db *dynamodb.Client
}

func NewRegisterAccountRepository(db *dynamodb.Client) itfdb.IBnFtRegisterAccountRepository {
	return &registerAccountRepository{db: db}
}
