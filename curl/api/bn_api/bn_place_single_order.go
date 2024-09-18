package api

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
	bnserivcemodelreq "tradething/app/bn/bn_future/bnservice_request_model"
	"tradething/config"
	"tradething/curl/api"
	"tradething/curl/curlcommon"
	"tradething/curl/curlservice"
)

type IBNApi interface {
	bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest |
		bnserivcemodelreq.SetLeverageBinanceServiceRequest |
		bnserivcemodelreq.QueryOrderBinanceServiceRequest
}

type placeSignleOrderRequest[T IBNApi] struct {
	bn_request       *T
	user_input_field []string
	user_input_kind  []reflect.Kind
	config           *config.AppConfig
}

func NewplaceSignleOrderRequest[T IBNApi](config *config.AppConfig) (api.IApi, error) {
	c := &placeSignleOrderRequest[T]{}
	c.bn_request = new(T)
	c.SetUserInputField()
	kind, err := curlcommon.GetFieldKind(c.bn_request)
	if err != nil {
		return nil, err
	}
	c.user_input_kind = kind
	c.config = config
	return c, nil
}

func (c *placeSignleOrderRequest[T]) GetUserInputField() []string {
	return c.user_input_field
}

func (c *placeSignleOrderRequest[T]) SetUserInputField() {
	fields := curlcommon.GetJsonTag(c.bn_request)
	c.user_input_field = fields
}

func (c *placeSignleOrderRequest[T]) SetUserInputValue(user_input_value []string) error {
	for idx, v := range user_input_value {
		k := c.user_input_kind[idx]
		f := c.user_input_field[idx]
		if f == "timestamp" {
			v = strconv.FormatInt(time.Now().Unix()*1000, 10)
		}
		if f == "type" && v == "" {
			v = "MARKET"
		}
		curlcommon.SetStructField(c.bn_request, f, v, k)
	}
	return nil
}

func (c *placeSignleOrderRequest[T]) ExecuteCurl() error {
	curl := curlservice.NewBinanceCurl(c.config)

	curl.SetMethod(http.MethodPost)
	base_url := c.config.BinanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1
	end_point := c.config.BinanceFutureUrl.SingleOrder
	url := fmt.Sprintf("%v%v", base_url, end_point)
	curl.SetUrl(url)

	urlValue := curlservice.SetBinanceSignUrl(c.bn_request, c.config.Secrets.BinanceSecretKey)
	curl.SetBody(urlValue)

	curl.PrepareCurl()

	err := curl.ExecuteCurl()
	if err != nil {
		return err
	}
	return nil
}
