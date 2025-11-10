package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtOpeningPosition struct {
	TableName       string `table:"bn_future_opening_position"`
	SymbolAccountId string `dynamodbav:"symbol_account_id" dynamodb:"symbol_account_id"` // primary key
	PositionSide    string `dynamodbav:"position_side" dynamodb:"position_side"`         // second index
	ClientId        string `dynamodbav:"client_id" dynamodb:"client_id"`
	Side            string `dynamodbav:"side" dynamodb:"side"`
	AmountB         string `dynamodbav:"amount_b" dynamodb:"amount_b"`
	CreatedAt       string `dynamodbav:"created_at" dynamodb:"created_at"`
}

func NewBnFtOpeningPositionTable() *BnFtOpeningPosition {
	return &BnFtOpeningPosition{}
}

func (b *BnFtOpeningPosition) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtOpeningPosition) SetSymbolAccountIdKey(symbol string, accountId string) {
	b.SymbolAccountId = fmt.Sprintf("%s_acc%s", strings.ToUpper(symbol), accountId)
}

func (b *BnFtOpeningPosition) GetAccountId() string {
	return strings.Split(b.SymbolAccountId, "_acc")[1]
}

func (b *BnFtOpeningPosition) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol_account_id": &types.AttributeValueMemberS{Value: b.SymbolAccountId},
		"position_side":     &types.AttributeValueMemberS{Value: b.PositionSide},
	}
}

func (b *BnFtOpeningPosition) GetPositionSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "position_side", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetClientIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "client_id", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "side", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetAmountBField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "amount_b", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetCreatedAtField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "created_at", "dynamodb")
	return v, t
}
