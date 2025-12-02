package db

import (
	"context"
	"tradething/app/bn/future/accumulation/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *BnFtAccumulationRepository) Get(ctx context.Context, clientId string) (*models.BnFtAccumulation, error) {
	table := models.NewBnFtAccumulationTable()
	table.ClientID = clientId
	result, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	if len(result.Item) == 0 {
		return nil, nil
	}

	res := &models.BnFtAccumulation{}
	err = attributevalue.UnmarshalMap(result.Item, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
