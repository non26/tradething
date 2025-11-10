package models

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtAdvancedPosition struct {
	TableName    string `table:"bn_future_advanced_position"`
	ClientID     string `dynamodbav:"client_id" dynamodb:"client_id"`
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	Side         string `dynamodbav:"side" dynamodb:"side"`
	AmountB      string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	AccountId    string `dynamodbav:"account_id" dynamodb:"account_id"`
}

func NewBnFtAdvancedPositionTable() *BnFtAdvancedPosition {
	return &BnFtAdvancedPosition{}
}

func (b *BnFtAdvancedPosition) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtAdvancedPosition) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientID},
	}
}

func (b *BnFtAdvancedPosition) GetSymbolField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "symbol", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPosition) GetPositionSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "position_side", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPosition) GetSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "side", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPosition) GetAmountBField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "amount_b", "dynamodb")
	return v, t
}

func (b *BnFtAdvancedPosition) GetAccountIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "account_id", "dynamodb")
	return v, t
}
