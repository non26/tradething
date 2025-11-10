package db

import (
	"context"
	"tradething/app/bn/future/position/domain"
	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	bntime "github.com/non26/tradepkg/pkg/bn/bn_time"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *openingPositionRepository) Upsert(ctx context.Context, position *domain.Position) error {
	table := models.NewBnFtOpeningPositionTable()
	table.SetSymbolAccountIdKey(position.Symbol, position.AccountId)
	table.PositionSide = position.PositionSide
	table.Side = position.Side
	table.AmountB = position.AmountB
	table.ClientId = position.ClientId
	table.CreatedAt = bntime.GetDBTime()

	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetSideField, table.Side)
	update_config.Set(table.GetAmountBField, table.AmountB)
	update_config.Set(table.GetClientIdField, table.ClientId)
	update_config.Set(table.GetCreatedAtField, table.CreatedAt)
	expression := update_config.BuildExpression()
	_, err := r.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table.GetTableName()),
		Key:                       table.GetKey(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	if err != nil {
		return err
	}
	return nil
}
