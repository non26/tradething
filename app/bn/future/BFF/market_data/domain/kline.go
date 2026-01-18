package domain

type Kline struct {
	Symbol           string
	Interval         string
	StartTime        int64
	EndTime          int64
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
