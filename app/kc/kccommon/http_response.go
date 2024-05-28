package kccommon

import (
	"fmt"
	"net/http"
)

// T response, k Request
type httpKcResponse[T any, K any] struct {
	res         *http.Response
	req         *httpKcRequest[K]
	transport   *http.Transport
	client      *http.Client
	body        T
	serviceName string
	err         error
}

func NewHttpKcResponse[T any, K any](
	req *httpKcRequest[K],
	serviceName string,
) *httpKcResponse[T, K] {
	return &httpKcResponse[T, K]{
		req:         req,
		serviceName: serviceName,
	}
}

func (h *httpKcResponse[T, K]) SetTransport() *httpKcResponse[T, K] {
	h.transport = &http.Transport{}
	return h
}

func (h *httpKcResponse[T, K]) SetClient() *httpKcResponse[T, K] {
	h.client = &http.Client{
		Transport: h.transport,
	}
	return h
}

func (h *httpKcResponse[T, K]) GetKcResponse() *httpKcResponse[T, K] {
	var err error
	h.res, err = h.client.Do(h.req.req)
	if err != nil {
		h.err = fmt.Errorf("%v-%v Response Error: %v", h.serviceName, h.req.urlPath, err.Error())
		h.res = nil
	}
	return h
}

func (h *httpKcResponse[T, K]) Error() (*httpKcResponse[T, K], error) {
	return h, h.err
}

func (h *httpKcResponse[T, K]) GetBody() T {
	return h.body
}

func (h *httpKcResponse[T, K]) GetKcResponseBody() *httpKcResponse[T, K] {
	defer h.res.Body.Close()
	kucoinResponse := NewKucoinResponse[T]()
	err := kucoinResponse.DecodeKcResponse(h.res.Body)
	if err != nil {
		h.err = fmt.Errorf("%v-%v Decode Error: %v", h.serviceName, h.req.urlPath, err.Error())
		return h
	}

	if h.res.StatusCode != 200 {
		h.err = fmt.Errorf("%v-%v KC code %v Error with HTTP code %v", h.serviceName, h.req.urlPath, kucoinResponse.GetKcInternalCode(), h.res.Status)
	}
	body := kucoinResponse.GetKcResult()
	h.body = body
	return h
}
