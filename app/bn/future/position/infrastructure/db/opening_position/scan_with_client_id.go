package db

import (
	"context"

	"tradething/app/bn/future/position/infrastructure/db/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (r *openingPositionRepository) ScanWithClientId(ctx context.Context, clientId string) (*models.BnFtOpeningPosition, error) {
	var err error
	var response *dynamodb.ScanOutput
	result := models.BnFtOpeningPosition{}
	table := models.NewBnFtOpeningPositionTable()
	table.ClientId = clientId
	response, err = r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table.GetTableName()),
		// Optional: Add a filter expression
		FilterExpression: aws.String("contains(#client_id, :value)"),
		ExpressionAttributeNames: map[string]string{
			"#client_id": "client_id", // Field name in DynamoDB table
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: clientId}, // Filter condition
		},
	})
	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, nil
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
