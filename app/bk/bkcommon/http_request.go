package bkcommon

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
	"tradething/common"
)

type httpBitkubRequest[K any] struct {
	req         *http.Request
	serviceName string
	method      string
	baseUrl     string
	urlPath     string
	payload     K
	err         error
}

func NewHttpBitkubRequest[K any](
	serviceName string,
	method string,
	baseUrl string,
	urlPath string,
	payload K,
) *httpBitkubRequest[K] {
	return &httpBitkubRequest[K]{
		serviceName: serviceName,
		method:      method,
		baseUrl:     baseUrl,
		urlPath:     urlPath,
		payload:     payload,
		err:         nil,
	}
}

func (h *httpBitkubRequest[K]) GetBkNewRequest() *httpBitkubRequest[K] {
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

func (h *httpBitkubRequest[K]) SetBkHeaders(apiKey string, sk string) *httpBitkubRequest[K] {
	var signature string
	t := time.Now().Unix()
	jsonPayload := common.ToJson(h.payload)
	if h.method == http.MethodPost {
		signature = CreateSignUrlMethod(
			http.MethodPost,
			t,
			h.urlPath,
			jsonPayload,
			sk,
			nil,
		)
	} else {
		signature = CreateSignUrlMethod(
			http.MethodGet,
			t,
			h.urlPath,
			nil,
			sk,
			common.JsonToMapString(jsonPayload),
		)
	}
	CreateBitkubHeaders(
		h.req,
		apiKey,
		t,
		signature,
	)
	return h
}

func (h *httpBitkubRequest[K]) GetHttpRequest() *http.Request {
	return h.req
}

func (h *httpBitkubRequest[K]) Error() (*httpBitkubRequest[K], error) {
	return h, h.err
}
