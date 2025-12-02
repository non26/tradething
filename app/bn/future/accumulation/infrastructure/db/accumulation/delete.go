package db

import (
	"context"
	"tradething/app/bn/future/accumulation/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *BnFtAccumulationRepository) Delete(ctx context.Context, clientId string) error {
	table := models.NewBnFtAccumulationTable()
	table.ClientID = clientId
	_, err := r.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table.GetTableName()),
		Key:       table.GetKey(),
	})
	return err
}
