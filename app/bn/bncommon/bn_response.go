package bncommon

import (
	"encoding/json"
	"errors"
	"net/http"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/common"
)

type binanceServiceHttpResponse[R any] struct {
	res   *http.Response
	bnres *R
}

func NewBinanceServiceHttpResponse[R any](res *http.Response) *binanceServiceHttpResponse[R] {
	b := binanceServiceHttpResponse[R]{
		res: res,
	}
	return &b
}

func (b *binanceServiceHttpResponse[R]) DecodeBinanceServiceResponse(
	binanceFutureServiceName string,
) error {
	if b.res.StatusCode != http.StatusOK {
		bnResponseError := new(bnservicemodelres.ResponseBinanceFutureError)
		json.NewDecoder(b.res.Body).Decode(bnResponseError)
		msg := common.FormatMessageOtherThanHttpStatus200(
			binanceFutureServiceName,
			b.res.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		defer b.res.Body.Close()
		return errors.New(msg)
	}

	bnResponse := new(R)
	json.NewDecoder(b.res.Body).Decode(bnResponse)
	b.bnres = bnResponse
	return nil
}

func (b *binanceServiceHttpResponse[R]) GetBinanceServiceResponse() *R {
	return b.bnres
}
