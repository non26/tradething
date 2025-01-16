package bnfuture

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	bntradereq "tradething/app/bn/bn_future/bnservice_request_model/trade"
	valueobject "tradething/app/bn/bn_future/value_object"

	mkreq "tradething/app/bn/bn_future/bnservice_request_model/market_data"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
)

type Position struct {
	positionSide  string
	side          string
	entryQuantity string // amountQ
	symbol        string
	leverageLevel int
	clientOrderId string
	// leverage      int
	stopLoss   *valueobject.StopLoss
	takeProfit *valueobject.TakeProfit
}

func (p *Position) GetPositionSide() string {
	return p.positionSide
}

func (p *Position) SetPositionSide(positionSide string) {
	p.positionSide = strings.ToUpper(positionSide)
}

func (p *Position) GetSide() string {
	return p.side
}

func (p *Position) SetSide(side string) {
	p.side = strings.ToUpper(side)
}

func (p *Position) GetEntryQuantity() string {
	return p.entryQuantity
}

func (p *Position) SetEntryQuantity(entryQuantity string) {
	p.entryQuantity = entryQuantity
}

func (p *Position) GetSymbol() string {
	return p.symbol
}

func (p *Position) SetSymbol(symbol string) {
	p.symbol = symbol
}

func (p *Position) GetLeverageLevel() int {
	return p.leverageLevel
}

func (p *Position) SetLeverageLevel(leverageLevel int) {
	p.leverageLevel = leverageLevel
}

func (p *Position) GetClientOrderId() string {
	return p.clientOrderId
}

func (p *Position) SetClientOrderId(clientOrderId string) {
	p.clientOrderId = clientOrderId
}

// func (p *Position) GetWatching() *valueobject.Watching {
// 	return p.watching
// }

// func (p *Position) SetWatching(watching *valueobject.Watching) {
// 	p.watching = watching
// }

// func (p *Position) IsWatchingNil() bool {
// 	return p.watching == nil
// }

// func (p *Position) IsWatchingStopLossNil() bool {
// 	return p.watching.StopLoss == nil
// }

// func (p *Position) IsWatchingTakeProfitNil() bool {
// 	return p.watching.TakeProfit == nil
// }

// func (p *Position) GetWatchingStopLoss() *valueobject.StopLoss {
// 	return p.watching.StopLoss
// }

// func (p *Position) GetWatchingTakeProfit() *valueobject.TakeProfit {
// 	return p.watching.TakeProfit
// }

func (p *Position) SetStopLoss(stopLoss *valueobject.StopLoss) {
	p.stopLoss = stopLoss
}

func (p *Position) GetStopLoss() *valueobject.StopLoss {
	return p.stopLoss
}

func (p *Position) SetTakeProfit(takeProfit *valueobject.TakeProfit) {
	p.takeProfit = takeProfit
}

func (p *Position) GetTakeProfit() *valueobject.TakeProfit {
	return p.takeProfit
}

func (p *Position) GetAmountQ() string {
	return p.entryQuantity
}

func (p *Position) IsStopLossNil() bool {
	return p.stopLoss == nil
}

func (p *Position) IsTakeProfitNil() bool {
	return p.takeProfit == nil
}

func (p *Position) IsLongPosition() bool {
	return p.positionSide == positionconstant.LONG
}

func (p *Position) IsShortPosition() bool {
	return p.positionSide == positionconstant.SHORT
}

func (p *Position) IsBuyOrder() bool {
	if p.IsLongPosition() && p.side == positionconstant.BUY {
		return true
	}
	if p.IsShortPosition() && p.side == positionconstant.SELL {
		return true
	}
	return false
}

func (p *Position) IsSellOrder() bool {
	if p.IsLongPosition() && p.side == positionconstant.SELL {
		return true
	}
	if p.IsShortPosition() && p.side == positionconstant.BUY {
		return true
	}
	return false
}

func (p *Position) AddEntryQuantity(entryQuantity string) {
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

func (p *Position) ToBinanceServiceModel() *bntradereq.PlaceSignleOrderBinanceServiceRequest {
	m := bntradereq.PlaceSignleOrderBinanceServiceRequest{
		PositionSide:  p.positionSide,
		Side:          p.side,
		EntryQuantity: p.entryQuantity,
		Symbol:        p.symbol,
		ClientOrderId: p.clientOrderId,
	}
	return &m
}

func (p *Position) ToBinanceFutureOpeningPositionRepositoryModel() *dynamodbmodel.BnFtOpeningPosition {
	m := dynamodbmodel.BnFtOpeningPosition{
		Symbol:       p.symbol,
		PositionSide: p.positionSide,
		AmountQ:      p.entryQuantity,
		Leverage:     fmt.Sprintf("%v", p.leverageLevel),
		ClientId:     p.clientOrderId,
		Side:         p.side,
		AmountB:      "",
	}
	watchingConfig := valueobject.Watching{
		StopLoss:   p.stopLoss,
		TakeProfit: p.takeProfit,
	}
	_json, err := json.Marshal(watchingConfig)
	if err != nil {
		return nil
	}
	m.SetWatchingConfig(_json)
	return &m
}

func (p *Position) ToBnCandleStickModel(starttime int64, endtime int64) *mkreq.CandleStickRequest {
	m := mkreq.CandleStickRequest{
		Symbol:    p.symbol,
		Interval:  p.stopLoss.Interval,
		StartTime: starttime,
		EndTime:   endtime,
	}
	return &m
}

func (p *Position) ToBnPositionHistoryRepositoryModel() *dynamodbmodel.BnFtHistory {
	m := dynamodbmodel.BnFtHistory{
		Symbol:       p.symbol,
		PositionSide: p.positionSide,
		ClientId:     p.clientOrderId,
	}
	return &m
}
