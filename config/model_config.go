package config

type AppConfig struct {
	Env              string           `mapstructure:"environment" json:"environment"`
	Port             string           `mapstructure:"port" json:"port"`
	ServiceName      ServiceName      `mapstructure:"service-name" json:"service-name"`
	Secrets          Secrets          `mapstructure:"secrets" json:"secrets"`
	BinanceFutureUrl BinanceFutureUrl `mapstructure:"binance-future-url" json:"binance-future-url"`
	OkxFutureUrl     OkxFutureUrl     `mapstructure:"okx-future-url" json:"okx-future-url"`
	KCFutureUrl      KCFutureUrl      `mapstructure:"kucoin-future-url" json:"kucoin-future-url"`
	KCSpotUrl        KCSpotUrl        `mapstructure:"kucoin-spot-url" json:"kucoin-spot-url"`
	KubSpotUrl       KubSpotUrl       `mapstructure:"kub-spot-url" json:"kub-spot-url"`
	Dynamodb         Dynamodb         `mapstructure:"dynamodb" json:"dynamodb"`
}

type Dynamodb struct {
	Region   string `mapstructure:"region" json:"region"`
	Ak       string `mapstructure:"ak" json:"ak"`
	Sk       string `mapstructure:"sk" json:"sk"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

type Secrets struct {
	BinanceApiKey       string `mapstructure:"binance-apiKey" json:"binance-apiKey"`
	BinanceSecretKey    string `mapstructure:"binance-secretKey" json:"binance-secretKey"`
	OkxApiKey           string `mapstructure:"okx-apiKey" json:"okx-apiKey"`
	OkxSecretKey        string `mapstructure:"okx-secretKey" json:"okx-secretKey"`
	OkxPassPhase        string `mapstructure:"okx-passPhase" json:"okx-passPhase"`
	BitkubApikey        string `mapstructure:"bitkub-apikey" json:"bitkub-apikey"`
	BitkubSecretKey     string `mapstructure:"bitkub-secretKey" json:"bitkub-secretKey"`
	KucoinApiKey        string `mapstructure:"kucoin-apiKey" json:"kucoin-apiKey"`
	KucoinApiKeyVersion string `mapstructure:"kucoin-apiKey-Version" json:"kucoin-apiKey-Version"`
	KucoinSecretKey     string `mapstructure:"kucoin-secretKey" json:"kucoin-secretKey"`
	KucoinPassphase     string `mapstructure:"kucoin-passphase" json:"kucoin-passphase"`
}

type ServiceName struct {
	BinanceFuture string `mapstructure:"binance-future" json:"binance-future"`
	BinanceSpot   string `mapstructure:"binance-spot" json:"binance-spot"`
	OKXFuture     string `mapstructure:"okx-future" json:"okx-future"`
	OKXSpot       string `mapstructure:"okx-spot" json:"okx-spot"`
	KucoinFuture  string `mapstructure:"kucoin-future" json:"kucoin-future"`
	KucoinSpot    string `mapstructure:"kucoin-spot" json:"kucoin-spot"`
	KubSpot       string `mapstructure:"kub-spot-url" json:"kub-spot-url"`
}

type BinanceFutureBaseUrl struct {
	BianceUrl1 string `mapstructure:"binance1" json:"binance1"`
	// BianceUrl2 string `mapstructure:"binance2"`
	// BianceUrl3 string `mapstructure:"binance3"`
	// BianceUrl4 string `mapstructure:"binance4"`
}

type BinanceFutureUrl struct {
	SetLeverage          string               `mapstructure:"set-leverage" json:"set-leverage"`
	SingleOrder          string               `mapstructure:"single-order" json:"single-order"`
	MultipleOrder        string               `mapstructure:"miltiple-order" json:"miltiple-order"`
	QueryOrder           string               `mapstructure:"query-order" json:"query-order"`
	ExchangeInfo         string               `mapstructure:"exchange-info" json:"exchange-info"`
	BinanceFutureBaseUrl BinanceFutureBaseUrl `mapstructure:"binance-future-baseUrl" json:"binance-future-baseUrl"`
}

type OkxFutureBaseUrl struct {
	Okx1 string `mapstructure:"okx1" json:""`
}

type OkxFutureUrl struct {
	SetLeverage        string           `mapstructure:"set-leverage" json:""`
	PlaceAPosition     string           `mapstructure:"place-position" json:""`
	PlaceMultiPosition string           `mapstructure:"multi-position" json:""`
	OkxFutureBaseUrl   OkxFutureBaseUrl `mapstructure:"okx-future-baseUrl" json:""`
}

type KCFutureUrl struct {
	BaseUrl       string `mapstructure:"base-url" json:"base-url"`
	PlaceOrderUrl string `mapstructure:"place-order-url" json:"place-order-url"`
}
type KCSpotUrl struct {
	BaseUrl       string `mapstructure:"base-url" json:"base-url"`
	PlaceOrderUrl string `mapstructure:"place-order-url" json:"place-order-url"`
}
type KubSpotUrl struct {
	BaseUrl       string `mapstructure:"base-url" json:"base-url"`
	SellOrderUrl  string `mapstructure:"sell-order-url" json:"sell-order-url"`
	BuyOrderUrl   string `mapstructure:"buy-order-url" json:"buy-order-url"`
	ServerTimeUrl string `mapstructure:"server-time-url" json:"server-time-url"`
	Balances      string `mapstructure:"balance-url" json:"balance-url"`
}
