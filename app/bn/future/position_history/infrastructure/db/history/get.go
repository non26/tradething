package db

import (
	"context"
	"tradething/app/bn/future/position_history/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *bnFtHistoryRepository) Get(ctx context.Context, clientId string) (*models.BnFtHistory, error) {
	table := models.NewBnFtHistoryTable()
	table.ClientId = clientId
	response, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKeyClientId(),
	})
	if err != nil {
		return nil, err
	}

	if response.Item == nil {
		return nil, nil
	}
	if len(response.Item) == 0 {
		return nil, nil
	}

	result := &models.BnFtHistory{}
	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
