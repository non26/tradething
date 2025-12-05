package models

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtAccumulation struct {
	AccumID      string `dynamodbav:"accum_id" dynamodb:"accum_id"`
	ClientID     string `dynamodbav:"client_id" dynamodb:"client_id"`
	MaxAccum     string `dynamodbav:"max_accum" dynamodb:"max_accum"`
	PresentAccum string `dynamodbav:"present_accum" dynamodb:"present_accum"`
}

func NewBnFtAccumulationTable() *BnFtAccumulation {
	return &BnFtAccumulation{}
}

func (b *BnFtAccumulation) GetTableName() string {
	return "bn_future_accumulation"
}

func (b *BnFtAccumulation) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientID},
	}
}

func (b *BnFtAccumulation) GetClientIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientID", "dynamodb")
	return v, t
}

func (b *BnFtAccumulation) GetMaxAccumField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "MaxAccum", "dynamodb")
	return v, t
}

func (b *BnFtAccumulation) GetPresentAccumField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "PresentAccum", "dynamodb")
	return v, t
}

func (b *BnFtAccumulation) GetAccumIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccumID", "dynamodb")
	return v, t
}
