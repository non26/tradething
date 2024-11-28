package bnfuture

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	repomodel "tradething/app/bn/bn_future/repository_model"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
// 	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
// 	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
// )

// func (d *dynamoDBRepository) GetQouteUSDT(ctx context.Context, symbol string) (*repomodel.BinanceFutureQouteUSDT, error) {
// 	var err error
// 	var response *dynamodb.GetItemOutput
// 	result := &repomodel.BinanceFutureQouteUSDT{}
// 	table := repomodel.NewBinanceFutureQouteUSTDTable()
// 	table.Symbol = symbol
// 	response, err = d.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
// 		TableName: aws.String(table.GetTableName()),
// 		Key:       table.GetKeyBySymbol(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = attributevalue.UnmarshalMap(response.Item, &result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (d *dynamoDBRepository) UpdateCountingSymbolQouteUSDT(ctx context.Context, qouteUSDT *repomodel.BinanceFutureQouteUSDT) error {
// 	table := repomodel.NewBinanceFutureQouteUSTDTable()
// 	table.BinanceFutureQouteUSDT = qouteUSDT
// 	input := &dynamodb.UpdateItemInput{
// 		TableName:        aws.String(table.GetTableName()),
// 		Key:              table.GetKeyBySymbol(),
// 		UpdateExpression: aws.String(fmt.Sprintf("set %v = :counting", table.GetCountingSymbolTableField())),
// 		ExpressionAttributeValues: map[string]types.AttributeValue{
// 			":counting": &types.AttributeValueMemberS{Value: table.GetNextCounting().String()},
// 		},
// 	}
// 	_, err := d.dynamodb.UpdateItem(ctx, input)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (d *dynamoDBRepository) InsertNewSymbolUSDT(ctx context.Context, symbol string) error {
// 	table := repomodel.NewBinanceFutureQouteUSTDTable()
// 	table.SetCounting(0)
// 	table.SetSymbol(symbol)
// 	item, err := attributevalue.MarshalMap(table.BinanceFutureQouteUSDT)
// 	if err != nil {
// 		log.Fatalf("Got error marshalling new movie item: %s", err)
// 	}
// 	input := &dynamodb.PutItemInput{
// 		TableName: aws.String(table.GetTableName()),
// 		Item:      item,
// 	}
// 	_, err = d.dynamodb.PutItem(ctx, input)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
