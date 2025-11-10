package db

import (
	"context"
	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *openingPositionRepository) Get(ctx context.Context, symbol string, accountId string, positionSide string) (*models.BnFtOpeningPosition, error) {
	table := models.NewBnFtOpeningPositionTable()
	table.PositionSide = positionSide
	table.SetSymbolAccountIdKey(symbol, accountId)
	response, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
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

	result := models.BnFtOpeningPosition{}
	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
