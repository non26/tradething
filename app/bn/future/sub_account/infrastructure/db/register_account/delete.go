package db

import (
	"context"
	models "tradething/app/bn/future/sub_account/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *registerAccountRepository) Delete(ctx context.Context, accountId string) error {
	table := models.NewBnFtRegisterAccountTable()
	table.AccountId = accountId
	_, err := r.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyAccountId(),
	})
	if err != nil {
		return err
	}
	return nil
}
