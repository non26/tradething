package db

import (
	"context"
	models "tradething/app/bn/future/sub_account/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *registerAccountRepository) Get(ctx context.Context, accountId string) (*models.BnFtRegisterAccount, error) {
	table := models.NewBnFtRegisterAccountTable()
	table.AccountId = accountId
	result, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyAccountId(),
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	sub_account := &models.BnFtRegisterAccount{}
	err = attributevalue.UnmarshalMap(result.Item, sub_account)
	if err != nil {
		return nil, err
	}
	return sub_account, nil
}
