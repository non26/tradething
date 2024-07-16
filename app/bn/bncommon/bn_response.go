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

	bnResponse := new(R)
	err := json.NewDecoder(b.res.Body).Decode(bnResponse)
	if err != nil {
		return err
	}
	b.bnres = bnResponse

	return nil
}

func (b *binanceServiceHttpResponse[R]) GetBinanceServiceResponse() *R {
	return b.bnres
}
