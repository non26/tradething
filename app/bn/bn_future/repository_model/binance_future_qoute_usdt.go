package bnfuture

// import (
// 	"reflect"
// 	"strconv"
// 	"tradething/app/bn/bncommon"

// 	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
// )

// type BinanceFutureQouteUSDT struct {
// 	Symbol         string `dynamodb:"symbol" dynamodbav:"symbol"` // primary key
// 	CountingSymbol int    `dynamodb:"counting_symbol" dynamodbav:"counting_symbol"`
// }

// type counting int

// func (c counting) Int() int {
// 	return int(c)
// }

// func (c counting) String() string {
// 	return strconv.Itoa(c.Int())
// }

// func newBinanceFutureQouteUSDT() *BinanceFutureQouteUSDT {
// 	return &BinanceFutureQouteUSDT{}
// }

// func (b *BinanceFutureQouteUSDT) GetNextCounting() counting {
// 	return counting(b.CountingSymbol + 1)
// }

// func (b *BinanceFutureQouteUSDT) SetCounting(counting int) {
// 	b.CountingSymbol = counting
// }

// func (b *BinanceFutureQouteUSDT) GetSymbol() string {
// 	return b.Symbol
// }

// func (b *BinanceFutureQouteUSDT) GetCounting() int {
// 	return b.CountingSymbol
// }

// func (b *BinanceFutureQouteUSDT) SetSymbol(symbol string) {
// 	b.Symbol = symbol
// }

// func (b *BinanceFutureQouteUSDT) IsExist() bool {
// 	return b.Symbol != ""
// }

// type BinanceFutureQouteUSTDTable struct {
// 	TableName string `table:"bn_future_qoute_usdt"`
// 	*BinanceFutureQouteUSDT
// }

// func (b *BinanceFutureQouteUSTDTable) GetTableName() string {
// 	return bncommon.GetStructTagValueByIndex(reflect.TypeOf(b).Elem(), "table", 0)
// }

// func NewBinanceFutureQouteUSTDTable() *BinanceFutureQouteUSTDTable {
// 	return &BinanceFutureQouteUSTDTable{
// 		BinanceFutureQouteUSDT: newBinanceFutureQouteUSDT(),
// 	}
// }

// func (b *BinanceFutureQouteUSTDTable) GetKeyBySymbol() map[string]types.AttributeValue {
// 	return map[string]types.AttributeValue{
// 		"symbol": &types.AttributeValueMemberS{Value: b.Symbol},
// 	}
// }

// func (b *BinanceFutureQouteUSTDTable) GetSymbolTableField() string {
// 	v, _ := bncommon.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Symbol", "dynamodb")
// 	return v
// }

// func (b *BinanceFutureQouteUSTDTable) GetCountingSymbolTableField() string {
// 	v, _ := bncommon.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CountingSymbol", "dynamodb")
// 	return v
// }
