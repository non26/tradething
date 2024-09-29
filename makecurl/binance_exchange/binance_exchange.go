package binanceexchange

import (
	binancefuture "tradething/makecurl/binance_exchange/binance_future"
	registorapi "tradething/makecurl/registor_api"
)

func RegistorBinanceFuture() registorapi.IRegistorApi {
	_registor_api := registorapi.NewRegistorApi()
	_registor_api.RegistorNewApi(binancefuture.NewOrder())

	return _registor_api
}
