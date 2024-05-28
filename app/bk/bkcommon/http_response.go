package bkcommon

import (
	"fmt"
	"net/http"
)

// T response
// K request
type httpBitkubResponse[T any, K any] struct {
	res         *http.Response
	req         *httpBitkubRequest[K]
	transport   *http.Transport
	client      *http.Client
	body        T
	serviceName string
	err         error
}

func NewHttpBitkubResponse[T any, K any](
	req *httpBitkubRequest[K],
	serviceName string,
) *httpBitkubResponse[T, K] {
	return &httpBitkubResponse[T, K]{
		req:         req,
		serviceName: serviceName,
	}
}

func (h *httpBitkubResponse[T, K]) SetTransport() *httpBitkubResponse[T, K] {
	h.transport = &http.Transport{}
	return h
}

func (h *httpBitkubResponse[T, K]) SetClient() *httpBitkubResponse[T, K] {
	h.client = &http.Client{
		Transport: h.transport,
	}
	return h
}

func (h *httpBitkubResponse[T, K]) GetNewBkResponse() *httpBitkubResponse[T, K] {
	var err error
	h.res, err = h.client.Do(h.req.req)
	if err != nil {
		h.err = fmt.Errorf("%v-%v Response Error: %v", h.serviceName, h.req.urlPath, err.Error())
		h.res = nil
	}
	return h
}

func (h *httpBitkubResponse[T, K]) Error() (*httpBitkubResponse[T, K], error) {
	return h, h.err
}

func (h *httpBitkubResponse[T, K]) GetBkResponseBody() *httpBitkubResponse[T, K] {
	if h.res.StatusCode == 429 {
		h.err = fmt.Errorf("%v-%v Rate Limit 429 Error", h.serviceName, h.req.urlPath)
		return h
	}
	defer h.res.Body.Close()
	bitkubResponse := NewBitKubResponse[T]()
	err := bitkubResponse.DecodeBkResponse(h.res.Body)
	if err != nil {
		h.err = fmt.Errorf("%v-%v Decode Error: %v", h.serviceName, h.req.urlPath, err.Error())
		return h
	}

	if h.res.StatusCode != 200 {
		h.err = fmt.Errorf("%v-%v BK code %v Error", h.serviceName, h.req.urlPath, bitkubResponse.GetBkInternalCode())
	}
	body := bitkubResponse.GetBkResult()
	h.body = body
	return h
}

func (h *httpBitkubResponse[T, K]) GetBody() T {
	return h.body
}
