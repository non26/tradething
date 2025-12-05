package models

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnFtRegisterAccount struct {
	AccountId   string `dynamodbav:"account_id" dynamodb:"account_id"`
	AccountName string `dynamodbav:"account_name" dynamodb:"account_name"`
	StartDate   string `dynamodbav:"start_date" dynamodb:"start_date"`
	EndDate     string `dynamodbav:"end_date" dynamodb:"end_date"`
}

func NewBnFtRegisterAccountTable() *BnFtRegisterAccount {
	return &BnFtRegisterAccount{}
}

func (b *BnFtRegisterAccount) GetTableName() string {
	return "bn_future_registor_account"
}

func (b *BnFtRegisterAccount) GetKeyAccountId() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"account_id": &types.AttributeValueMemberS{Value: b.AccountId},
	}
}

func (b *BnFtRegisterAccount) GetAccountIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccountId", "dynamodb")
	return v, t
}

func (b *BnFtRegisterAccount) GetAccountNameField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccountName", "dynamodb")
	return v, t
}

func (b *BnFtRegisterAccount) GetStartDateField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "StartDate", "dynamodb")
	return v, t
}

func (b *BnFtRegisterAccount) GetEndDateField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "EndDate", "dynamodb")
	return v, t
}
