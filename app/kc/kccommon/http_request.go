package kccommon

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
	"tradething/common"
)

type httpKcRequest[K any] struct {
	req         *http.Request
	serviceName string
	method      string
	baseUrl     string
	urlPath     string
	payload     K
	err         error
}

func NewHttpKcRequest[K any](
	serviceName string,
	method string,
	baseUrl string,
	urlPath string,
	payload K,
) *httpKcRequest[K] {
	return &httpKcRequest[K]{
		serviceName: serviceName,
		method:      method,
		baseUrl:     baseUrl,
		urlPath:     urlPath,
		payload:     payload,
		err:         nil,
	}
}

func (h *httpKcRequest[K]) GetKcNewRequest() *httpKcRequest[K] {
	var req *http.Request
	var err error
	_url := h.baseUrl + h.urlPath
	if h.method == http.MethodPost {
		req, err = http.NewRequest(http.MethodPost, _url, bytes.NewReader(common.ToJson(h.payload)))
	} else {
		req, err = http.NewRequest(http.MethodGet, _url, nil)
	}
	if err != nil {
		h.err = fmt.Errorf("%v-%v Request Error: %v", h.serviceName, h.urlPath, err.Error())
		h.req = nil
	}
	h.req = req
	return h
}

func (h *httpKcRequest[K]) SetKcHeaders(
	apiKey string,
	apiKeyVersion string,
	passphase string,
	sk string) *httpKcRequest[K] {
	jsonString := common.JsonToString(common.ToJson(h.payload))
	kcHeader := KucoinHeaderInfo{
		Request:    h.req,
		Method:     h.method,
		UrlPath:    h.urlPath,
		Passphase:  passphase,
		SecretKey:  sk,
		ApiKey:     apiKey,
		KucoinTime: time.Now().Unix(),
	}
	kcHeader.AddHeaders(jsonString)
	return h
}

func (h *httpKcRequest[K]) GetHttpRequest() *http.Request {
	return h.req
}

func (h *httpKcRequest[K]) Error() (*httpKcRequest[K], error) {
	return h, h.err
}
