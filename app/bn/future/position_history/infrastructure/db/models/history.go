package models

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtHistory struct {
	TableName    string `table:"bn_future_history"`
	ClientId     string `dynamodbav:"client_id" dynamodb:"client_id"` // primary key
	Symbol       string `dynamodbav:"symbol" dynamodb:"symbol"`
	PositionSide string `dynamodbav:"position_side" dynamodb:"position_side"`
	AccountId    string `dynamodbav:"account_id" dynamodb:"account_id"`
}

func NewBnFtHistoryTable() *BnFtHistory {
	return &BnFtHistory{}
}

func (b *BnFtHistory) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtHistory) GetKeyClientId() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

func (b *BnFtHistory) GetSymbolField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "symbol", "dynamodb")
	return v, t
}

func (b *BnFtHistory) GetPositionSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "position_side", "dynamodb")
	return v, t
}

func (b *BnFtHistory) GetAccountIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "account_id", "dynamodb")
	return v, t
}
