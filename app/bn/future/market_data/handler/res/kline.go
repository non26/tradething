package res

import "tradething/app/bn/future/market_data/domain"

type KlinesResponse struct {
	KlineData []*KlineResponse
}

func (k *KlinesResponse) FromDomain(domainKlines []*domain.Kline) *KlinesResponse {
	for _, domainKline := range domainKlines {
		klineResponse := &KlineResponse{}
		k.KlineData = append(k.KlineData, klineResponse.FromDomain(domainKline))
	}
	return k
}

type KlineResponse struct {
	Open                  string `json:"open"`
	High                  string `json:"high"`
	Low                   string `json:"low"`
	Close                 string `json:"close"`
	IsCloseHigherThanOpen bool   `json:"is_green_candle"`
}

func (k *KlineResponse) FromDomain(domainKline *domain.Kline) *KlineResponse {
	return &KlineResponse{
		Open:                  domainKline.Open,
		High:                  domainKline.High,
		Low:                   domainKline.Low,
		Close:                 domainKline.Close,
		IsCloseHigherThanOpen: domainKline.IsGreenCandle,
	}
}
