package bnservicemodelres

import bnhandlerres "tradething/app/bn/app/model/handlermodel/future/response"

type ExchangeInfoRateLimit struct {
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type ExchangeInfoAssets struct {
	Asset             string `json:"asset"`
	MarginAvailable   bool   `json:"marginAvailable"`
	AutoAssetExchange int    `json:"autoAssetExchange"`
}

type ExchangeInfoSymbolsFilter struct {
	FilterType        string `json:"filterType"`
	MaxPrice          string `json:"maxPrice,omitempty"`
	MinPrice          string `json:"minPrice,omitempty"`
	TickSize          string `json:"tickSize,omitempty"`
	MaxQty            string `json:"maxQty,omitempty"`
	MinQty            string `json:"minQty,omitempty"`
	StepSize          string `json:"stepSize,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Notional          string `json:"notional,omitempty"`
	MultiplierUp      string `json:"multiplierUp,omitempty"`
	MultiplierDown    string `json:"multiplierDown,omitempty"`
	MultiplierDecimal int    `json:"multiplierDecimal,omitempty"`
}

type ExchangeInfoSymbols struct {
	Symbol                string                      `json:"symbol"`
	Pair                  string                      `json:"pair"`
	ContractType          string                      `json:"contractType"`
	DeliveryDate          int64                       `json:"deliveryDate"`
	OnboardDate           int64                       `json:"onboardDate"`
	Status                string                      `json:"status"`
	MaintMarginPercent    string                      `json:"maintMarginPercent"`
	RequiredMarginPercent string                      `json:"requiredMarginPercent"`
	BaseAsset             string                      `json:"baseAsset"`
	QuoteAsset            string                      `json:"quoteAsset"`
	MarginAsset           string                      `json:"marginAsset"`
	PricePrecision        int                         `json:"pricePrecision"`
	QuantityPrecision     int                         `json:"quantityPrecision"`
	BaseAssetPrecision    int                         `json:"baseAssetPrecision"`
	QuotePrecision        int                         `json:"quotePrecision"`
	UnderlyingType        string                      `json:"underlyingType"`
	UnderlyingSubType     []string                    `json:"underlyingSubType"`
	SettlePlan            int                         `json:"settlePlan"`
	TriggerProtect        string                      `json:"triggerProtect"`
	Filters               []ExchangeInfoSymbolsFilter `json:"filters"`
	OrderType             []string                    `json:"OrderType"`
	TimeInForce           []string                    `json:"timeInForce"`
	LiquidationFee        string                      `json:"liquidationFee"`
	MarketTakeBound       string                      `json:"marketTakeBound"`
}

type ExchangeInfoResponse struct {
	ExchangeFilters []any                   `json:"exchangeFilters"`
	RateLimits      []ExchangeInfoRateLimit `json:"rateLimits"`
	ServerTime      int64                   `json:"serverTime"`
	Assets          []ExchangeInfoAssets    `json:"assets"`
	Symbols         []ExchangeInfoSymbols   `json:"symbols"`
	Timezone        string                  `json:"timezone"`
}

func (e *ExchangeInfoResponse) ToHandlerResponse() bnhandlerres.ExchangeInfoHandlerResponse {
	m := bnhandlerres.ExchangeInfoHandlerResponse{}
	for _, symbol := range e.Symbols {
		m.Symbols = append(m.Symbols, symbol.Symbol)
	}
	return m
}
