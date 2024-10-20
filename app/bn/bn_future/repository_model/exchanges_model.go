package bnfuture

type Exchange struct {
	Id           int    `dynamodbav:"id"`
	ExchangeName string `dynamodbav:"exchange_name"`
}

type ExchangeTable struct {
	TableName string `table:"exchanges"`
	*Exchange
}
