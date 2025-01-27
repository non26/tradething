package bnfuture

import (
	"fmt"
	"strconv"
	"strings"
	bntradereq "tradething/app/bn/bn_future/bnservice_request/trade"
	valueobject "tradething/app/bn/bn_future/value_object"

	mkreq "tradething/app/bn/bn_future/bnservice_request/market_data"

	positionconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
	"github.com/shopspring/decimal"
)

type Position struct {
	positionSide  string
	side          string
	entryQuantity string // amountB
	symbol        string
	leverageLevel int
	clientOrderId string
	stopLoss      *valueobject.StopLoss
	takeProfit    *valueobject.TakeProfit
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

func (p *Position) GetClientId() string {
	return p.clientOrderId
}

func (p *Position) SetClientId(clientOrderId string) {
	p.clientOrderId = clientOrderId
}

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

func (p *Position) GetAmountB() string {
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

func (p *Position) AddMoreAmountB(amountB string) error {
	amountQInt, err := decimal.NewFromString(amountB)
	if err != nil {
		return err
	}
	prevAmountQInt, err := decimal.NewFromString(p.entryQuantity)
	if err != nil {
		return err
	}
	p.entryQuantity = amountQInt.Add(prevAmountQInt).String()
	return nil
}

func (p *Position) ToBinanceServiceModel() *bntradereq.PlacePosition {
	m := bntradereq.PlacePosition{
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
		AmountB:      p.entryQuantity,
		ClientId:     p.clientOrderId,
		Side:         p.side,
	}
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
