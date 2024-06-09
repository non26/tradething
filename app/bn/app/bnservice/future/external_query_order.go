package bnservice

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	model "tradething/app/bn/app/model/handlermodel/future"
	"tradething/app/bn/bncommon"
)

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *model.QueryOrderBinanceHandlerRequest,
) (*http.Response, error) {
	endPoint := bfes.binanceFutureUrl.QueryOrder
	_url := fmt.Sprintf("%v%v", bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, endPoint)

	t := time.Now().Unix() * 1000
	tt := strconv.FormatInt(t, 10)
	data := url.Values{}
	data.Set("symbol", request.Symbol)
	data.Set("timestamp", tt)
	data.Set("origClientOrderId", request.Symbol)
	encodeData := bncommon.CreateBinanceSignature(&data, bfes.secrets.BinanceSecretKey)

	req, err := http.NewRequest(http.MethodGet, _url, nil)
	// req, err := http.NewRequest(http.MethodGet, _url, strings.NewReader(encodeData))
	if err != nil {
		return nil, errors.New("Query Order Request Error: " + err.Error())
	}

	// q := req.URL.Query()
	// q.Add("symbol", request.Symbol)
	// q.Add("timestamp", tt)
	// q.Add("signature", encodeData)
	// req.URL.RawQuery = q.Encode()
	req.URL.RawQuery = encodeData
	println(req.URL.RawQuery)
	println(req.URL.Path)
	println(req.URL.Host)

	req.Header.Add("X-MBX-APIKEY", bfes.secrets.BinanceApiKey)
	req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")

	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return nil, errors.New("Query Order Response Error: " + err.Error())
	}

	return res, nil
}
