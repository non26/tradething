package config

type AppConfig struct {
	Env         string      `mapstructure:"environment" json:"environment"`
	Port        string      `mapstructure:"port" json:"port"`
	ServiceName ServiceName `mapstructure:"service-name" json:"service-name"`
	// Secrets          Secrets          `mapstructure:"secrets" json:"secrets"`
	// BinanceFutureUrl BinanceFutureUrl `mapstructure:"binance-future-url" json:"binance-future-url"`
	Dynamodb Dynamodb `mapstructure:"dynamodb" json:"dynamodb"`
	// BinanceSpotUrl   BinanceSpotUrl   `mapstructure:"binance-spot-url" json:"binance-spot-url"`
	BinanceAdaptorFutureUsdt BinanceAdaptorFutureUsdt `mapstructure:"binance-adaptor-future-usdt" json:"binance-adaptor-future-usdt"`
}

func (a AppConfig) GetPortWithFormat() string {
	return ":" + a.Port
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

type ServiceName struct {
	BinanceFuture string `mapstructure:"binance-future" json:"binance-future"`
	BinanceSpot   string `mapstructure:"binance-spot" json:"binance-spot"`
}

type BinanceAdaptorFutureUsdt struct {
	BaseUrl           string `mapstructure:"base-url" json:"base-url"`
	NewOrder          string `mapstructure:"new-order" json:"new-order"`
	KlinesCandleStick string `mapstructure:"klines-candle-stick" json:"klines-candle-stick"`
}
