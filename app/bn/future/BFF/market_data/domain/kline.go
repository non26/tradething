package domain

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
