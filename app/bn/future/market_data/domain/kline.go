package domain

import (
	"strings"
	"tradething/app/bn/future/market_data/infrastructure/adaptor/market_data/req"

	bntime "github.com/non26/tradepkg/pkg/bn/bn_time"
)

type Kline struct {
	Symbol           string
	Interval         string
	StartTime        string
	EndTime          string
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	CloseTime        int64
	QuoteVolume      string
	NumberOfTrades   int64
	BuyerBaseVolume  string
	BuyerQuoteVolume string
	IsGreenCandle    bool
}

func (k *Kline) GetPreviousStartAndEndTimeInUnixTimestamp() (int64, int64) {
	interval := strings.ToLower(k.Interval)
	timeframe := interval[len(interval)-1:]
	unit := interval[:len(interval)-1]
	switch timeframe {
	case "h":
		startTime, endTime := bntime.GetBinanceStartAndEndTimeInHourTimeFrame(unit)
		previousStartTime, previousEndTime := bntime.GetBinancePreviousStartAndEndTimeInHourTimeFrame(startTime, endTime, unit)
		return bntime.GetSpecificBnTimestamp(&previousStartTime), bntime.GetSpecificBnTimestamp(&previousEndTime)
	case "d":
		startTime, endTime := bntime.GetBinanceStartAndEndTimeInDayTimeFrame(unit)
		previousStartTime, previousEndTime := bntime.GetBinancePreviousStartAndEndTimeInDayTimeFrame(startTime, endTime, unit)
		return bntime.GetSpecificBnTimestamp(&previousStartTime), bntime.GetSpecificBnTimestamp(&previousEndTime)
	default:
		return 0, 0
	}
}

func (k *Kline) ToMarketDataAdaptorKlineRequest() req.KlineRequest {
	r := req.KlineRequest{
		Symbol:    k.Symbol,
		Interval:  k.Interval,
		StartTime: k.StartTime,
		EndTime:   k.EndTime,
	}
	return r
}
