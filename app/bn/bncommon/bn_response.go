package bncommon

import (
	"encoding/json"
	"errors"
	"net/http"
	bnservicemodelres "tradething/app/bn/bn_future/bnservice_response_model"
	"tradething/common"
)

type IBinanceServiceHttpResponse[R any] interface {
	SetResponse(res *http.Response)
	DecodeBinanceServiceResponse(binanceFutureServiceName string) error
	GetBinanceServiceResponse() *R
}

type binanceServiceHttpResponse[R any] struct {
	res   *http.Response
	bnres *R
}

func NewBinanceServiceHttpResponse[R any]() *binanceServiceHttpResponse[R] {
	b := binanceServiceHttpResponse[R]{}
	return &b
}

func (b *binanceServiceHttpResponse[R]) SetResponse(res *http.Response) {
	b.res = res
}

func (b *binanceServiceHttpResponse[R]) DecodeBinanceServiceResponse(
	binanceFutureServiceName string,
) error {
	defer b.res.Body.Close()
	if b.res.StatusCode != http.StatusOK {
		bnResponseError := new(bnservicemodelres.ResponseBinanceFutureError)
		json.NewDecoder(b.res.Body).Decode(bnResponseError)
		if bnResponseError.Code == -2013 {
			b.bnres = new(R)
			return nil
		}
		msg := common.FormatMessageOtherThanHttpStatus200(
			binanceFutureServiceName,
			b.res.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		return errors.New(msg)
	}
	b.bnres = new(R)
	switch any(*b.bnres).(type) {
	case struct{}:
	default:
		bnResponse := new(R)
		err := json.NewDecoder(b.res.Body).Decode(bnResponse)
		if err != nil {
			return err
		}
		b.bnres = bnResponse
	}

	return nil
}

func (b *binanceServiceHttpResponse[R]) GetBinanceServiceResponse() *R {
	return b.bnres
}
