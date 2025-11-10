package db

import (
	"context"
	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *advancedPositionRepository) Delete(ctx context.Context, clientId string) error {
	table := models.NewBnFtAdvancedPositionTable()
	table.ClientID = clientId
	_, err := r.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	if err != nil {
		return err
	}
	return nil
}
