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
Todo:

	deal with the request without response
*/
func CallBinance[Q, P any](
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

	bnreq := bncommon.NewBinanceServiceHttpRequest[Q]()
	err := bnreq.NewBinanceHttpRequest(_url)
	if err != nil {
		return nil, err
	}

	data := request.GetData().(*Q)
	signature := bnreq.CreateRequestSignUrl(data, secret_key)
	switch method {
	case http.MethodPost:
		bnreq.RequestPostMethod(signature)
	default:
		bnreq.RequestGetMethod(signature)
	}
	bnreq.AddHeader(api_key)

	bntransport := bncommon.NewBinanceTransport(&http.Transport{})
	bnclient := bncommon.NewBinanceSerivceHttpClient(bntransport.GetTransport())
	err = bnclient.Do(bnreq.GetBinanceRequest())
	if err != nil {
		return nil, err
	}

	bnres := bncommon.NewBinanceServiceHttpResponse[P](
		bnclient.GetBinanceHttpClientResponse())
	err = bnres.DecodeBinanceServiceResponse(service_name)
	if err != nil {
		return nil, err
	}

	return bnres.GetBinanceServiceResponse(), nil
}
