package db

import (
	"context"
	"tradething/app/bn/future/position/domain"
	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *advancedPositionRepository) Upsert(ctx context.Context, position *domain.Position) error {
	table := models.NewBnFtAdvancedPositionTable()
	table.ClientID = position.ClientId
	table.Symbol = position.Symbol
	table.PositionSide = position.PositionSide
	table.Side = position.Side
	table.AmountB = position.AmountB
	table.AccountId = position.AccountId

	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetSymbolField, position.Symbol)
	update_config.Set(table.GetPositionSideField, position.PositionSide)
	update_config.Set(table.GetSideField, position.Side)
	update_config.Set(table.GetAmountBField, position.AmountB)
	update_config.Set(table.GetAccountIdField, position.AccountId)
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
