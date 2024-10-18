package bnfuture

type BinanceFutureOpeningPositionTable struct {
	Id         string `dynamodbav:"id"`
	ClientId   string `dynamodbav:"client_id"`
	ExchangeId int    `dynamodbav:"exchange_id"`
	Symbol     string `dyanmodbav:"symbol"`
	// Type               string `dynamodbav:"type"`
	Leverage int `dynamodbav:"leverage"`
	Position int `dynamodbav:"position"`
	// Amount             string `dynamodbav:"amount"`
	// AmountCurrency     string `dynamodbav:"amount_currency"`
	Amount_Q string `dynamodbav:"amount_q"`
	Amount_B string `dynamodbav:"amount_b"`
	// IsBuyFilled        bool   `dynamodbav:"is_buy_filled"`
	// IsSellFilled       bool   `dynamodbav:"is_sell_filled"`
	// IsActive           bool   `dynamodbav:"is_active"`
	// BuyOrderCreatedAt  string `dynamodbav:"buy_order_created_at"`
	// SellOrderCreatedAt string `dynamodbav:"sell_order_created_at"`
	// BuyUpdatedAt       string `dynamodbav:"buy_updated_at"`
	// SellUpdatedAt      string `dynamodbav:"sell_updated_at"`
	BuyOrderCreatedAt  string `dynamodbav:"buy_created_at"`
	SellOrderCreatedAt string `dynamodbav:"sell_created_at"`
}
