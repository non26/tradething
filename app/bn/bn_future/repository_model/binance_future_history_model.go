package bnfuture

type BinanceFutureHistoryTable struct {
	Id                 string `dynamodbav:"id"`
	ExchangeId         int    `dynamodbav:"exchange_id"`
	ClientId           string `dynamodbav:"client_id"`
	Pnl                string `dynamodb:"pnl"`
	CreatedAt          string `dynamodbav:"created_at"`
	BuyOrderCreatedAt  string `dynamodbav:"buy_order_created_at"`
	SellOrderCreatedAt string `dynamodbav:"sell_order_created_at"`
	BuyUpdatedAt       string `dynamodb:"buy_updated_at"`
	SellUpdatedAt      string `dynamodbav:"sell_updated_at"`
}
