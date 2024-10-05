package bnfuture

type ExchangeTable struct {
	Id           int    `dynamodbav:"id"`
	ExchangeName string `dynamodbav:"exchange_name"`
}
