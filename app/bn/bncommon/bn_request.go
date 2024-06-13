package bncommon

import (
	"io"
	"net/http"
	"strings"
)

/*
T is type of request
*/
type binanceServiceHttpRequest[T any] struct {
	req *http.Request
}

func NewBinanceServiceHttpRequest[T any]() *binanceServiceHttpRequest[T] {
	return &binanceServiceHttpRequest[T]{}
}

func (b *binanceServiceHttpRequest[T]) NewBinanceHttpRequest(
	url string,
) error {
	req, err := http.NewRequest(
		"", url, nil,
	)
	if err != nil {
		return err
	}
	b.req = req
	return nil
}

func (b *binanceServiceHttpRequest[T]) CreateRequestSignUrl(request *T, secretKey string) string {
	data := GetQueryStringFromStructType(*request)
	sig := CreateBinanceSignature(data, secretKey)
	return sig
}

func (b *binanceServiceHttpRequest[T]) RequestPostMethod(signature string) {
	b.req.Method = http.MethodPost
	var body io.Reader = strings.NewReader(signature)
	rc, ok := body.(io.ReadCloser)
	if !ok {
		rc = io.NopCloser(body)
	}
	b.req.Body = rc
}

func (b *binanceServiceHttpRequest[T]) RequestGetMethod(signature string) {
	b.req.Method = http.MethodGet
	b.req.URL.RawQuery = signature
}

func (b *binanceServiceHttpRequest[T]) AddHeader(apiKey string) {
	b.req.Header.Add("X-MBX-APIKEY", apiKey)
	b.req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")
}

func (b *binanceServiceHttpRequest[T]) GetBinanceRequest() *http.Request {
	return b.req
}
