package models

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtAccumulation struct {
	TableName    string `table:"bn_future_accumulation"`
	AccumID      string `dynamodbav:"accum_id" dynamodb:"accum_id"`
	ClientID     string `dynamodbav:"client_id" dynamodb:"client_id"`
	MaxAccum     string `dynamodbav:"max_accum" dynamodb:"max_accum"`
	PresentAccum string `dynamodbav:"present_accum" dynamodb:"present_accum"`
}

func NewBnFtAccumulationTable() *BnFtAccumulation {
	return &BnFtAccumulation{}
}

func (b *BnFtAccumulation) GetTableName() string {
	return utils.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
}

func (b *BnFtAccumulation) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.AccumID},
	}
}

func (b *BnFtAccumulation) GetClientIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "client_id", "dynamodb")
	return v, t
}

func (b *BnFtAccumulation) GetMaxAccumField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "max_accum", "dynamodb")
	return v, t
}

func (b *BnFtAccumulation) GetPresentAccumField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "present_accum", "dynamodb")
	return v, t
}

func (b *BnFtAccumulation) GetAccumIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "accum_id", "dynamodb")
	return v, t
}
