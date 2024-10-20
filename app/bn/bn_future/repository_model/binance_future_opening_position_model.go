package bnfuture

import (
	"reflect"
	"tradething/app/bn/bncommon"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BinanceFutureOpeningPosition struct {
	Id         string `dynamodbav:"id"`
	ClientId   string `dynamodbav:"client_id"`
	ExchangeId int    `dynamodbav:"exchange_id"`
	Symbol     string `dyanmodbav:"symbol"`
	// Type               string `dynamodbav:"type"`
	Leverage     int    `dynamodbav:"leverage"`
	PositionSide string `dynamodbav:"position_side"`
	// Amount             string `dynamodbav:"amount"`
	// AmountCurrency     string `dynamodbav:"amount_currency"`
	AmountQ string `dynamodbav:"amount_q"`
	AmountB string `dynamodbav:"amount_b"`
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

func (b *BinanceFutureOpeningPosition) GetKeyBySymbol() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
	}
}

func (b *BinanceFutureOpeningPosition) GetKeyByClientID() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"client_id": &types.AttributeValueMemberS{Value: b.ClientId},
	}
}

func newBinanceFutureOpeningPosition() *BinanceFutureOpeningPosition {
	return &BinanceFutureOpeningPosition{
		ExchangeId: 1,
	}
}

type BinanceFutureOpeningPositionTable struct {
	TableName string `table:"bn_future_opening_position_table"`
	*BinanceFutureOpeningPosition
}

func NewBinanceFutureOpeningPositionTable() *BinanceFutureOpeningPositionTable {
	return &BinanceFutureOpeningPositionTable{
		BinanceFutureOpeningPosition: newBinanceFutureOpeningPosition(),
	}
}

func (b *BinanceFutureOpeningPositionTable) GetTableName() string {
	return bncommon.GetStructTagValueByIndex(reflect.TypeOf(b), "table", 0)
}
