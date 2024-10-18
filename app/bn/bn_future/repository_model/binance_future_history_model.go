package bnfuture

type BinanceFutureHistoryTable struct {
	Id            string `dynamodbav:"id"`
	ExchangeId    int    `dynamodbav:"exchange_id"`
	ClientId      string `dynamodbav:"client_id"`
	Pnl           string `dynamodb:"pnl"`
	Leverage      string `dynamodb:"leverage"`
	Position      string `dynamodb:"position"`
	Symbol        string `dynamodb:"symbol"`
	BuyCreatedAt  string `dynamodbav:"buy_created_at"`
	SellCreatedAt string `dynamodbav:"sell_created_at"`
	// CreatedAt          string `dynamodbav:"created_at"`
	// BuyOrderCreatedAt  string `dynamodbav:"buy_order_created_at"`
	// SellOrderCreatedAt string `dynamodbav:"sell_order_created_at"`
	// BuyUpdatedAt       string `dynamodb:"buy_updated_at"`
	// SellUpdatedAt      string `dynamodbav:"sell_updated_at"`
}
