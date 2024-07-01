package init

import (
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	"tradething/config"
	"tradething/curl/api"
	bnapi "tradething/curl/api/bn_api"
	"tradething/curl/constant"
)

func InitApiRequestBodyMenu() map[string]map[string]func(config *config.AppConfig) (api.IApi, error) {
	m := make(map[string]map[string]func(config *config.AppConfig) (api.IApi, error))
	bn := make(map[string]func(config *config.AppConfig) (api.IApi, error))
	bn[constant.Bn_CreateOrder] = bnapi.NewplaceSignleOrderRequest[bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest]
	bn[constant.Bn_QueryOrder] = bnapi.NewQueryOrderrequest[bnserivcemodelreq.QueryOrderBinanceServiceRequest]
	m[constant.Binance] = bn

	return m
}
