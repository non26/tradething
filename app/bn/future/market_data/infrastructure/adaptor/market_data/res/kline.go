package res

type KlineAdaptorResponse [][]interface{}

func (k *KlineAdaptorResponse) ToReader() *KlineReaderAdaptorResponse {
	klineReaderAdaptorResponse := &KlineReaderAdaptorResponse{
		KlineData: make([]KlineDataAdaptorResponse, len(*k)),
	}

	for i, v := range *k {
		klineReaderAdaptorResponse.KlineData[i] = KlineDataAdaptorResponse{
			OpenTime:         v[0].(float64),
			Open:             v[1].(string),
			High:             v[2].(string),
			Low:              v[3].(string),
			Close:            v[4].(string),
			Volume:           v[5].(string),
			CloseTime:        v[6].(float64),
			QuoteVolume:      v[7].(string),
			NumberOfTrades:   v[8].(float64),
			BuyerBaseVolume:  v[9].(string),
			BuyerQuoteVolume: v[10].(string),
		}
	}
	return klineReaderAdaptorResponse
}

type KlineReaderAdaptorResponse struct {
	KlineData []KlineDataAdaptorResponse
}

type KlineDataAdaptorResponse struct {
	OpenTime         float64
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	CloseTime        float64
	QuoteVolume      string
	NumberOfTrades   float64
	BuyerBaseVolume  string
	BuyerQuoteVolume string
}
