package bnservice

import (
	"fmt"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	"tradething/app/bn/bncommon"
)

/*
generic Q for requset model
generic P for response model
*/
type ICallBinance[Q, P any] interface {
	CallBinance(
		request bnserivcemodelreq.IBnFutureServiceRequest,
		base_url string,
		end_point string,
		method string,
		secret_key string,
		api_key string,
		service_name string,
	) (*P, error)
}

type callBinance[Q any, P any] struct {
	http_request   bncommon.IBinanceServiceHttpRequest[Q]
	http_response  bncommon.IBinanceServiceHttpResponse[P]
	http_transport bncommon.IBinanceServiceHttpTransport
	http_client    bncommon.IBinanceSerivceHttpClient
}

func NewCallBinance[Q, P any](
	http_request bncommon.IBinanceServiceHttpRequest[Q],
	http_response bncommon.IBinanceServiceHttpResponse[P],
	http_transport bncommon.IBinanceServiceHttpTransport,
	http_client bncommon.IBinanceSerivceHttpClient,
) ICallBinance[Q, P] {
	c := callBinance[Q, P]{
		http_request,
		http_response,
		http_transport,
		http_client,
	}
	return &c
}

func (c *callBinance[Q, P]) CallBinance(
	request bnserivcemodelreq.IBnFutureServiceRequest,
	base_url string,
	end_point string,
	method string,
	secret_key string,
	api_key string,
	service_name string,
) (*P, error) {
	_url := fmt.Sprintf("%v%v", base_url, end_point)
	request.PrepareRequest()

	err := c.http_request.NewBinanceHttpRequest(_url)
	if err != nil {
		return nil, err
	}

	data := request.GetData().(*Q)
	signature := c.http_request.CreateRequestSignUrl(data, secret_key)
	switch method {
	case http.MethodPost:
		c.http_request.RequestPostMethod(signature)
	default:
		c.http_request.RequestGetMethod(signature)
	}
	c.http_request.AddHeader(api_key)

	c.http_client.SetClient(c.http_transport.GetTransport())
	err = c.http_client.Do(c.http_request.GetBinanceRequest())
	if err != nil {
		return nil, err
	}

	c.http_response.SetResponse(c.http_client.GetBinanceHttpClientResponse())
	err = c.http_response.DecodeBinanceServiceResponse(service_name)
	if err != nil {
		return nil, err
	}

	return c.http_response.GetBinanceServiceResponse(), nil
}
