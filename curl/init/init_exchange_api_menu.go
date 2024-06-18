package init

import "tradething/curl/constant"

func InitExchangeApiMenu() map[string][]string {
	m := map[string][]string{}
	m[constant.Binance] = constant.BinanceApis
	return m
}
