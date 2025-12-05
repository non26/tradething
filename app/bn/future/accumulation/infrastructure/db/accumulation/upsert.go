package db

import (
	"context"
	"tradething/app/bn/future/accumulation/domain"
	"tradething/app/bn/future/accumulation/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *BnFtAccumulationRepository) Upsert(ctx context.Context, accumulation *domain.Accumulation) error {
	table := models.NewBnFtAccumulationTable()
	table.AccumID = accumulation.AccumID
	table.ClientID = accumulation.ClientID
	table.MaxAccum = accumulation.MaxAccum
	table.PresentAccum = accumulation.PresentAccum
	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetAccumIdField, accumulation.AccumID)
	update_config.Set(table.GetMaxAccumField, accumulation.MaxAccum)
	update_config.Set(table.GetPresentAccumField, accumulation.PresentAccum)
	_, err := r.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.GetTableName()),
		Key:                       table.GetKey(),
		UpdateExpression:          update_config.BuildExpression(),
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	return err
}
