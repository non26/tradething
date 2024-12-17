package bnfuture

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	valueobject "tradething/app/bn/bn_future/value_object"

	mkreq "tradething/app/bn/bn_future/bnservice_request_model/market_data"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

type PlaceSignleOrderServiceRequest struct {
	positionSide  string
	side          string
	entryQuantity string // amountQ
	symbol        string
	leverageLevel int
	clientOrderId string
	leverage      int
	stopLoss      *valueobject.StopLoss
	takeProfit    *valueobject.TakeProfit
}

func (p *PlaceSignleOrderServiceRequest) GetPositionSide() string {
	return p.positionSide
}

func (p *PlaceSignleOrderServiceRequest) SetPositionSide(positionSide string) {
	p.positionSide = strings.ToUpper(positionSide)
}

func (p *PlaceSignleOrderServiceRequest) GetSide() string {
	return p.side
}

func (p *PlaceSignleOrderServiceRequest) SetSide(side string) {
	p.side = strings.ToUpper(side)
}

func (p *PlaceSignleOrderServiceRequest) GetEntryQuantity() string {
	return p.entryQuantity
}

func (p *PlaceSignleOrderServiceRequest) SetEntryQuantity(entryQuantity string) {
	p.entryQuantity = entryQuantity
}

func (p *PlaceSignleOrderServiceRequest) GetSymbol() string {
	return p.symbol
}

func (p *PlaceSignleOrderServiceRequest) SetSymbol(symbol string) {
	p.symbol = symbol
}

func (p *PlaceSignleOrderServiceRequest) GetLeverageLevel() int {
	return p.leverageLevel
}

func (p *PlaceSignleOrderServiceRequest) SetLeverageLevel(leverageLevel int) {
	p.leverageLevel = leverageLevel
}

func (p *PlaceSignleOrderServiceRequest) GetClientOrderId() string {
	return p.clientOrderId
}

func (p *PlaceSignleOrderServiceRequest) SetClientOrderId(clientOrderId string) {
	p.clientOrderId = clientOrderId
}

// func (p *PlaceSignleOrderServiceRequest) GetWatching() *valueobject.Watching {
// 	return p.watching
// }

// func (p *PlaceSignleOrderServiceRequest) SetWatching(watching *valueobject.Watching) {
// 	p.watching = watching
// }

// func (p *PlaceSignleOrderServiceRequest) IsWatchingNil() bool {
// 	return p.watching == nil
// }

// func (p *PlaceSignleOrderServiceRequest) IsWatchingStopLossNil() bool {
// 	return p.watching.StopLoss == nil
// }

// func (p *PlaceSignleOrderServiceRequest) IsWatchingTakeProfitNil() bool {
// 	return p.watching.TakeProfit == nil
// }

// func (p *PlaceSignleOrderServiceRequest) GetWatchingStopLoss() *valueobject.StopLoss {
// 	return p.watching.StopLoss
// }

// func (p *PlaceSignleOrderServiceRequest) GetWatchingTakeProfit() *valueobject.TakeProfit {
// 	return p.watching.TakeProfit
// }

func (p *PlaceSignleOrderServiceRequest) SetStopLoss(stopLoss *valueobject.StopLoss) {
	p.stopLoss = stopLoss
}

func (p *PlaceSignleOrderServiceRequest) GetStopLoss() *valueobject.StopLoss {
	return p.stopLoss
}

func (p *PlaceSignleOrderServiceRequest) SetTakeProfit(takeProfit *valueobject.TakeProfit) {
	p.takeProfit = takeProfit
}

func (p *PlaceSignleOrderServiceRequest) GetTakeProfit() *valueobject.TakeProfit {
	return p.takeProfit
}

func (p *PlaceSignleOrderServiceRequest) GetAmountQ() string {
	return p.entryQuantity
}

func (p *PlaceSignleOrderServiceRequest) IsStopLossNil() bool {
	return p.stopLoss == nil
}

func (p *PlaceSignleOrderServiceRequest) IsTakeProfitNil() bool {
	return p.takeProfit == nil
}

func (p *PlaceSignleOrderServiceRequest) AddEntryQuantity(entryQuantity string) {
	var currentQuantity float64
	currentQuantity, err := strconv.ParseFloat(p.entryQuantity, 64)
	if err != nil {
		return
	}
	var additionalQuantity float64
	additionalQuantity, err = strconv.ParseFloat(entryQuantity, 64)
	if err != nil {
		return
	}
	p.entryQuantity = fmt.Sprintf("%v", currentQuantity+additionalQuantity)
}

func (p *PlaceSignleOrderServiceRequest) ToBinanceServiceModel() *bntradereq.PlaceSignleOrderBinanceServiceRequest {
	m := bntradereq.PlaceSignleOrderBinanceServiceRequest{
		PositionSide:  p.positionSide,
		Side:          p.side,
		EntryQuantity: p.entryQuantity,
		Symbol:        p.symbol,
		ClientOrderId: p.clientOrderId,
	}
	return &m
}

func (p *PlaceSignleOrderServiceRequest) ToBinanceFutureOpeningPositionRepositoryModel() *dynamodbmodel.BinanceFutureOpeningPosition {
	m := dynamodbmodel.BinanceFutureOpeningPosition{
		Symbol:             p.symbol,
		PositionSide:       p.positionSide,
		AmountQ:            p.entryQuantity,
		Leverage:           fmt.Sprintf("%v", p.leverageLevel),
		ClientId:           p.clientOrderId,
		Side:               p.side,
		AmountB:            "",
		BuyOrderCreatedAt:  time.Now().Format(time.DateTime),
		SellOrderCreatedAt: time.Now().Format(time.DateTime),
	}
	return &m
}

func (p *PlaceSignleOrderServiceRequest) ToBnCandleStickModel(starttime int64, endtime int64) *mkreq.CandleStickRequest {
	m := mkreq.CandleStickRequest{
		Symbol:    p.symbol,
		Interval:  p.stopLoss.Interval,
		StartTime: starttime,
		EndTime:   endtime,
	}
	return &m
}

func (p *PlaceSignleOrderServiceRequest) ToBnPositionHistoryRepositoryModel() *dynamodbmodel.BinanceFutureHistory {
	m := dynamodbmodel.BinanceFutureHistory{
		Symbol:       p.symbol,
		PositionSide: p.positionSide,
		ClientId:     p.clientOrderId,
	}
	return &m
}
