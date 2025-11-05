package db

import (
	"context"
	"tradething/app/bn/future/sub_account/domain"
	models "tradething/app/bn/future/sub_account/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *registerAccountRepository) Upsert(ctx context.Context, sub_account *domain.SubAccount) error {
	table := models.NewBnFtRegisterAccountTable()
	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetAccountIdField, sub_account.AccountId)
	update_config.Set(table.GetAccountNameField, sub_account.AccountName)
	update_config.Set(table.GetStartDateField, sub_account.StartDate)
	update_config.Set(table.GetEndDateField, sub_account.EndDate)
	_, err := r.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:        aws.String(table.GetTableName()),
		Key:              table.GetKeyAccountId(),
		UpdateExpression: update_config.BuildExpression(),
	})
	if err != nil {
		return err
	}
	return nil
}
