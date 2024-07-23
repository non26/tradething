package config

import "github.com/spf13/viper"

func ReadConfig(path_config string) (c *AppConfig, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&c)

	if c.Env == "local" {
		c.BinanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1 = "https://testnet.binancefuture.com"
		// c.Binance.SpotBaseUrl = "https://testnet.binance.vision"
	}

	return c, nil
}
