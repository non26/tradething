package db

import (
	"context"
	"tradething/app/bn/future/accumulation/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *BnFtAccumulationRepository) GetAll(ctx context.Context) ([]*models.BnFtAccumulation, error) {
	table := models.NewBnFtAccumulationTable()
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

	items := []*models.BnFtAccumulation{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
