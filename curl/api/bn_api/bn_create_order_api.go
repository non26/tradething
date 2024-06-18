package api

import (
	"fmt"
	"net/http"
	"reflect"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnhandlerreq "tradething/app/bn/app/model/handlermodel/future/request"
	"tradething/config"
	"tradething/curl/api"
	"tradething/curl/curlcommon"
	"tradething/curl/curlservice"
)

type createOrderRequest struct {
	user_input       *bnhandlerreq.PlaceSignleOrderHandlerRequest
	bn_request       *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest
	user_input_field []string
	user_input_kind  []reflect.Kind
	config           *config.AppConfig
}

func NewCreateOrderRequest(config *config.AppConfig) (api.IApi, error) {
	ureq := bnhandlerreq.PlaceSignleOrderHandlerRequest{}
	bnreq := bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest{}

	c := createOrderRequest{}
	c.user_input = &ureq
	c.bn_request = &bnreq
	c.SetUserInputField()
	kind, err := curlcommon.GetFieldKind(c.user_input)
	if err != nil {
		return nil, err
	}
	c.user_input_kind = kind
	c.config = config
	return &c, nil
}

func (c *createOrderRequest) GetUserInputField() []string {
	return c.user_input_field
}

func (c *createOrderRequest) SetUserInputField() {
	fields := curlcommon.GetJsonTag(c.user_input)
	c.user_input_field = fields
}

func (c *createOrderRequest) SetUserInputValue(user_input_value []string) error {
	for idx, v := range user_input_value {
		k := c.user_input_kind[idx]
		f := c.user_input_field[idx]
		curlcommon.SetStructField(c.user_input, f, v, k)
	}
	c.bn_request = c.user_input.ToBinanceServiceModel()
	return nil
}

func (c *createOrderRequest) GenerateCurl() error {
	curl := curlservice.NewBinanceCurl(c.config)

	curl.SetMethod(http.MethodPost)
	url := fmt.Sprintf("%v%v", c.config.BinanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, c.config.BinanceFutureUrl.SingleOrder)
	curl.SetUrl(url)

	c.bn_request.PrepareRequest()
	urlValue := curlservice.SetBinanceSignUrl(c.bn_request, c.config.Secrets.BinanceSecretKey)
	curl.SetBody(urlValue)

	curl.PrepareCurl()
	err := curl.ExecuteCurl()
	if err != nil {
		return err
	}
	return nil
}
