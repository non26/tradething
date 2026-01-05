package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	bntime "github.com/non26/tradepkg/pkg/bn/bn_time"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtOpeningPosition struct {
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
	return "bn_future_opening_position"
}

func (b *BnFtOpeningPosition) SetSymbolAccountIdKey(symbol string, accountId string) {
	b.SymbolAccountId = fmt.Sprintf("%s_acc%s", strings.ToUpper(symbol), accountId)
}

func (b *BnFtOpeningPosition) GetAccountId() string {
	return strings.Split(b.SymbolAccountId, "_acc")[1]
}

func (b *BnFtOpeningPosition) GetSymbol() string {
	return strings.Split(b.SymbolAccountId, "_acc")[0]
}

func (b *BnFtOpeningPosition) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol_account_id": &types.AttributeValueMemberS{Value: b.SymbolAccountId},
		"position_side":     &types.AttributeValueMemberS{Value: b.PositionSide},
	}
}

func (b *BnFtOpeningPosition) GetPositionSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PositionSide", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetClientIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientId", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetSideField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Side", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetAmountBField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AmountB", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) GetCreatedAtField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CreatedAt", "dynamodb")
	return v, t
}

func (b *BnFtOpeningPosition) SetCreatedAt() {
	b.CreatedAt = bntime.GetDBTime()
}
