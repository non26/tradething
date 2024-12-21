package bnfuture

import (
	"strings"
	bnsvcreq "tradething/app/bn/bn_future/bot_service_model"
)

type BotTimeframeExeIntervalHandlerRequest struct {
	BotId        string `json:"bot_id"`
	BotOrderID   string `json:"bot_order_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	// Config       BotTimeframeExeIntervalConfig `json:"config"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// type BotTimeframeExeIntervalConfig struct {
// 	Timeframe string  `json:"timeframe"`
// 	Interval  string  `json:"interval"`
// 	AmountQ   float64 `json:"amount_q"`
// 	StopLoss  *valueobject.StopLossConfig `json:"stop_loss"`
// }

func (b *BotTimeframeExeIntervalHandlerRequest) Validate() error {
	return nil
}

func (b *BotTimeframeExeIntervalHandlerRequest) Transform() error {
	b.StartDate = transformToRFC3339(b.StartDate)
	b.EndDate = transformToRFC3339(b.EndDate)
	return nil
}

func transformToRFC3339(_time string) string {
	date_time := strings.Split(_time, " ")
	date := date_time[0]
	time := date_time[1]
	date_time_utc := date + "T" + time + "Z+07:00"
	return date_time_utc
}

func (b *BotTimeframeExeIntervalHandlerRequest) ToBotServiceRequest() (*bnsvcreq.BotTimeframeExeIntervalRequest, error) {
	svcmodel := &bnsvcreq.BotTimeframeExeIntervalRequest{}
	svcmodel.SetBotId(b.BotId)
	svcmodel.SetBotOrderID(b.BotOrderID)
	svcmodel.SetSymbol(b.Symbol)
	svcmodel.SetPositionSide(b.PositionSide)
	svcmodel.SetStartDate(b.StartDate)
	svcmodel.SetEndDate(b.EndDate)
	return svcmodel, nil
}
