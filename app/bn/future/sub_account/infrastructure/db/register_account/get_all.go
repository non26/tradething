package db

import (
	"context"
	models "tradething/app/bn/future/sub_account/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *registerAccountRepository) GetAll(ctx context.Context) ([]*models.BnFtRegisterAccount, error) {
	table := models.NewBnFtRegisterAccountTable()
	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}
	if result.Items == nil {
		return nil, nil
	}
	if len(result.Items) == 0 {
		return nil, nil
	}
	sub_accounts := []*models.BnFtRegisterAccount{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &sub_accounts)
	if err != nil {
		return nil, err
	}
	return sub_accounts, nil
}
