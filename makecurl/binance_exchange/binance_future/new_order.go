package binancefuture

import (
	"net/http"
	apiinfo "tradething/makecurl/registor_api/api_info"
)

func NewOrder() apiinfo.IApiInfo {
	api_alias := "new order"
	api_endpoint := "/fapi/v1/order"
	method := http.MethodPost

	_api := apiinfo.NewApi(
		api_alias, api_endpoint, method,
	)
	_api.SetRequestBody("symbol", "string")
	_api.SetRequestBody("side", "string")
	_api.SetRequestBody("type", "string")
	_api.SetRequestBody("quantity", "float64")
	_api.SetRequestBody("timestamp", "string")
	_api.SetRequestBody("newClientOrderId", "string")
	return _api
}
