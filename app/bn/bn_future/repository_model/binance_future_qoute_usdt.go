package bnfuture

type BinanceFutureQouteUSDT struct {
	Id             string `dynamodbav:"id"`
	ExchangeId     int    `dynamodb:"exchange_id"`
	Symbol         string `dynamodb:"symbol"`
	CountingSymbol int    `dynamodbav:"counting_symbol"`
}

type BinanceFutureQouteUSTDTable struct {
	TableName string `table:"bn_future_qoute_usdt"`
	*BinanceFutureQouteUSDT
}
