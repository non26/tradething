package db

import (
	"context"
	"tradething/app/bn/future/position_history/domain"
	"tradething/app/bn/future/position_history/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *bnFtHistoryRepository) Insert(ctx context.Context, history *domain.BnFtHistory) error {
	table := models.NewBnFtHistoryTable()
	table.ClientId = history.ClientId
	table.Symbol = history.Symbol
	table.PositionSide = history.PositionSide
	table.CreatedAt = history.CreatedAt
	item, err := attributevalue.MarshalMap(table)
	if err != nil {
		return err
	}
	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table.GetTableName()),
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}
