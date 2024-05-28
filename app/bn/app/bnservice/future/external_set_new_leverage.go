package bnservice

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	model "tradetoolv2/app/bn/app/model/handlermodel/future"
	"tradetoolv2/app/bn/bncommon"
)

func (bfes *binanceFutureExternalService) SetNewLeverage(
	c context.Context,
	request *model.SetLeverageHandlerRequest) (*http.Response, error) {
	base := bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1
	_url := fmt.Sprintf("%v%v", base, bfes.binanceFutureUrl.SetLeverage)

	t := time.Now().Unix() * 1000
	data := url.Values{}
	data.Set("symbol", request.Symbol)
	data.Set("leverage", strconv.Itoa(request.Leverage))
	data.Set("timestamp", strconv.FormatInt(t, 10))
	encodeData := bncommon.CreateBinanceSignature(&data, bfes.secrets.BinanceSecretKey)

	req, err := http.NewRequest(
		http.MethodPost,
		_url,
		strings.NewReader(encodeData),
	)
	if err != nil {
		return nil, errors.New("SetNewLeverage Request Error: " + err.Error())
	}

	req.Header.Add("X-MBX-APIKEY", bfes.secrets.BinanceApiKey)
	req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")

	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return nil, errors.New("SetNewLeverage Response Error: " + err.Error())
	}

	return res, nil
}
