package config

type AppConfig struct {
	Env              string           `mapstructure:"environment" json:"environment"`
	Port             string           `mapstructure:"port" json:"port"`
	ServiceName      ServiceName      `mapstructure:"service-name" json:"service-name"`
	Secrets          Secrets          `mapstructure:"secrets" json:"secrets"`
	BinanceFutureUrl BinanceFutureUrl `mapstructure:"binance-future-url" json:"binance-future-url"`
	Dynamodb         Dynamodb         `mapstructure:"dynamodb" json:"dynamodb"`
	BinanceSpotUrl   BinanceSpotUrl   `mapstructure:"binance-spot-url" json:"binance-spot-url"`
}

func (c *AppConfig) IsLocal() bool {
	return c.Env == "local"
}

func (c *AppConfig) IsPrd() bool {
	return c.Env == "prd"
}

type Dynamodb struct {
	Region   string `mapstructure:"region" json:"region"`
	Ak       string `mapstructure:"ak" json:"ak"`
	Sk       string `mapstructure:"sk" json:"sk"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

type Secrets struct {
	BinanceApiKey              string `mapstructure:"binance-apiKey" json:"binance-apiKey"`
	BinanceSecretKey           string `mapstructure:"binance-secretKey" json:"binance-secretKey"`
	BinanceSpotApiKey          string `mapstructure:"binance-spot-apiKey" json:"binance-spot-apiKey"`
	BinanceSpotSecretKey       string `mapstructure:"binance-spot-secretKey" json:"binance-spot-secretKey"`
	BinanceSubAccountApikey    string `mapstructure:"binance-subAccount-apikey" json:"binance-subAccount-apikey"`
	BinanceSubAccountSecretKey string `mapstructure:"binance-subAccount-secretKey" json:"binance-subAccount-secretKey"`
}

type ServiceName struct {
	BinanceFuture string `mapstructure:"binance-future" json:"binance-future"`
	BinanceSpot   string `mapstructure:"binance-spot" json:"binance-spot"`
}

type BinanceFutureBaseUrl struct {
	BianceUrl1 string `mapstructure:"binance1" json:"binance1"`
}

type BinanceFutureMarketData struct {
	CandleStick string `mapstructure:"candle-stick" json:"candle-stick"`
}

type BinanceFutureUrl struct {
	SetLeverage             string                  `mapstructure:"set-leverage" json:"set-leverage"`
	SingleOrder             string                  `mapstructure:"single-order" json:"single-order"`
	MultipleOrder           string                  `mapstructure:"miltiple-order" json:"miltiple-order"`
	QueryOrder              string                  `mapstructure:"query-order" json:"query-order"`
	ExchangeInfo            string                  `mapstructure:"exchange-info" json:"exchange-info"`
	BinanceFutureBaseUrl    BinanceFutureBaseUrl    `mapstructure:"binance-future-baseUrl" json:"binance-future-baseUrl"`
	BinanceFutureMarketData BinanceFutureMarketData `mapstructure:"market-data" json:"market-data"`
}

type BinanceSpotUrl struct {
	SingleOrder        string               `mapstructure:"single-order" json:"single-order"`
	MultipleOrders     string               `mapstructure:"miltiple-order" json:"miltiple-order"`
	QueryOrder         string               `mapstructure:"query-order" json:"query-order"`
	ExcahngeInfo       string               `mapstructure:"exchange-info" json:"exchange-info"`
	BinanceSpotBaseUrl BinanceFutureBaseUrl `mapstructure:"binance-spot-baseUrl" json:"binance-spot-baseUrl"`
}

type BinanceSpotBaseUrl struct {
	BianceUrl1 string `mapstructure:"binance1" json:"binance1"`
}
