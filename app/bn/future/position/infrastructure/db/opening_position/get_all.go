package db

import (
	"context"
	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *openingPositionRepository) GetAll(ctx context.Context) ([]*models.BnFtOpeningPosition, error) {
	table := models.NewBnFtOpeningPositionTable()
	response, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
	})
	if err != nil {
		return nil, err
	}
	if response.Items == nil {
		return nil, nil
	}
	if len(response.Items) == 0 {
		return nil, nil
	}
	items := []*models.BnFtOpeningPosition{}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
