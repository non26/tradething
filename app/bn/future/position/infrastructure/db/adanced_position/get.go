package db

import (
	"context"
	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *advancedPositionRepository) Get(ctx context.Context, clientId string) (*models.BnFtAdvancedPosition, error) {
	table := models.NewBnFtAdvancedPositionTable()
	table.ClientID = clientId
	result := models.BnFtAdvancedPosition{}
	response, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	if err != nil {
		return nil, err
	}
	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
