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

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *model.PlaceSignleOrderHandlerRequest,
) (*http.Response, error) {
	endPoint := bfes.binanceFutureUrl.SingleOrder
	_url := fmt.Sprintf("%v%v", bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, endPoint)

	t := time.Now().Unix() * 1000
	data := url.Values{}
	data.Set("symbol", request.Symbol)
	data.Set("side", request.Side)
	data.Set("positionSide", request.PositionSide)
	data.Set("type", "MARKET")
	data.Set("quantity", fmt.Sprintf("%v", request.EntryQuantity))
	data.Set("timestamp", strconv.FormatInt(t, 10))
	encodeData := bncommon.CreateBinanceSignature(&data, bfes.secrets.BinanceSecretKey)

	req, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(encodeData))
	if err != nil {
		return nil, errors.New("PlaceSingleOrder Request Error: " + err.Error())
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
		return nil, errors.New("PlaceSingleOrder Response Error: " + err.Error())
	}

	return res, nil
}
