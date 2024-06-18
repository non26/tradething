package init

import (
	"tradething/config"
	iapi "tradething/curl/api"
	bnapi "tradething/curl/api/bn_api"
	"tradething/curl/constant"
)

func InitApiRequestBodyMenu() map[string]map[string]func(config *config.AppConfig) (iapi.IApi, error) {
	m := make(map[string]map[string]func(config *config.AppConfig) (iapi.IApi, error))
	bn := make(map[string]func(config *config.AppConfig) (iapi.IApi, error))
	bn[constant.Bn_CreateOrder] = bnapi.NewCreateOrderRequest
	m[constant.Binance] = bn
	return m
}
