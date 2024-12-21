package bnfuture

import (
	"strconv"
	"time"

	bnsvcreq "tradething/app/bn/bn_future/bnservice_request_model/trade"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func parseRFC3339ToUTC(_time string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, _time)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime.UTC(), nil
}

type BotTimeframeExeIntervalRequest struct {
	botId        string
	botOrderID   string
	symbol       string
	positionSide string
	timeframe    string
	interval     string
	amountQ      float64
	startDate    time.Time
	endDate      time.Time
	// StopLoss     *valueobject.StopLossConfig
}

func (b *BotTimeframeExeIntervalRequest) SetBotId(botId string) {
	b.botId = botId
}

func (b *BotTimeframeExeIntervalRequest) SetBotOrderID(botOrderID string) {
	b.botOrderID = botOrderID
}

func (b *BotTimeframeExeIntervalRequest) SetSymbol(symbol string) {
	b.symbol = symbol
}

func (b *BotTimeframeExeIntervalRequest) SetPositionSide(positionSide string) {
	b.positionSide = positionSide
}

func (b *BotTimeframeExeIntervalRequest) SetTimeframe(timeframe string) {
	b.timeframe = timeframe
}

func (b *BotTimeframeExeIntervalRequest) SetInterval(interval string) {
	b.interval = interval
}

func (b *BotTimeframeExeIntervalRequest) SetAmountQ(amountQ float64) {
	b.amountQ = amountQ
}

func (b *BotTimeframeExeIntervalRequest) SetStartDate(startDate string) error {
	parsedTime, err := parseRFC3339ToUTC(startDate)
	if err != nil {
		return err
	}
	b.startDate = parsedTime
	return nil
}

func (b *BotTimeframeExeIntervalRequest) SetEndDate(endDate string) error {
	parsedTime, err := parseRFC3339ToUTC(endDate)
	if err != nil {
		return err
	}
	b.endDate = parsedTime
	return nil
}

func (b *BotTimeframeExeIntervalRequest) GetStartDate() time.Time {
	return b.startDate
}

func (b *BotTimeframeExeIntervalRequest) GetEndDate() time.Time {
	return b.endDate
}

func (b *BotTimeframeExeIntervalRequest) GetBotId() string {
	return b.botId
}

func (b *BotTimeframeExeIntervalRequest) GetBotOrderID() string {
	return b.botOrderID
}

func (b *BotTimeframeExeIntervalRequest) GetSymbol() string {
	return b.symbol
}

func (b *BotTimeframeExeIntervalRequest) GetPositionSide() string {
	return b.positionSide
}

func (b *BotTimeframeExeIntervalRequest) GetTimeframe() string {
	return b.timeframe
}

func (b *BotTimeframeExeIntervalRequest) GetInterval() string {
	return b.interval
}

func (b *BotTimeframeExeIntervalRequest) GetAmountQ() float64 {
	return b.amountQ
}

func (b *BotTimeframeExeIntervalRequest) ToBnFtOpeningPosition() *dynamodbmodel.BnFtOpeningPosition {
	m := dynamodbmodel.BnFtOpeningPosition{
		Symbol:       b.symbol,
		PositionSide: b.positionSide,
		ClientId:     b.botOrderID,
	}
	return &m
}

func (b *BotTimeframeExeIntervalRequest) ToBnFtPlaceSingleOrderServiceRequest(side string, orderType string) *bnsvcreq.PlaceSignleOrderBinanceServiceRequest {
	m := bnsvcreq.PlaceSignleOrderBinanceServiceRequest{
		Symbol:        b.symbol,
		PositionSide:  b.positionSide,
		ClientOrderId: b.botOrderID,
		EntryQuantity: strconv.FormatFloat(b.amountQ, 'f', -1, 64),
		Side:          side,
		Type:          orderType,
	}
	return &m
}

func (b *BotTimeframeExeIntervalRequest) ToBnFtDeleteBotOnRun() *dynamodbmodel.BnFtBotOnRun {
	m := dynamodbmodel.BnFtBotOnRun{
		BotID:      b.botId,
		BotOrderID: b.botOrderID,
	}
	return &m
}

func (b *BotTimeframeExeIntervalRequest) ToBnFtHistory() *dynamodbmodel.BnFtHistory {
	m := dynamodbmodel.BnFtHistory{
		ClientId:     b.botOrderID,
		Symbol:       b.symbol,
		PositionSide: b.positionSide,
	}
	return &m
}

func (b *BotTimeframeExeIntervalRequest) ToBnFtBotOnRun() *dynamodbmodel.BnFtBotOnRun {
	m := dynamodbmodel.BnFtBotOnRun{
		BotID:        b.botId,
		BotOrderID:   b.botOrderID,
		Symbol:       b.symbol,
		PositionSide: b.positionSide,
	}
	return &m
}
