package bnservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/app/bn/bncommon"
	"tradething/common"
)

// func (bfes *binanceFutureExternalService) SetNewLeverage2(
// 	c context.Context,
// 	request *model.SetLeverageHandlerRequest) (*http.Response, error) {
// 	base := bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1
// 	_url := fmt.Sprintf("%v%v", base, bfes.binanceFutureUrl.SetLeverage)

// 	t := time.Now().Unix() * 1000
// 	data := url.Values{}
// 	data.Set("symbol", request.Symbol)
// 	data.Set("leverage", strconv.Itoa(request.Leverage))
// 	data.Set("timestamp", strconv.FormatInt(t, 10))
// 	encodeData := bncommon.CreateBinanceSignature(&data, bfes.secrets.BinanceSecretKey)

// 	req, err := http.NewRequest(
// 		http.MethodPost,
// 		_url,
// 		strings.NewReader(encodeData),
// 	)
// 	if err != nil {
// 		return nil, errors.New("SetNewLeverage Request Error: " + err.Error())
// 	}

// 	req.Header.Add("X-MBX-APIKEY", bfes.secrets.BinanceApiKey)
// 	req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")

// 	transport := &http.Transport{}
// 	client := &http.Client{
// 		Transport: transport,
// 	}
// 	var res *http.Response
// 	res, err = client.Do(req)
// 	if err != nil {
// 		return nil, errors.New("SetNewLeverage Response Error: " + err.Error())
// 	}

// 	return res, nil
// }

func (bfes *binanceFutureExternalService) SetNewLeverage(
	ctx context.Context,
	request *bnserivcemodelreq.SetLeverageBinanceServiceRequest) error {

	base := bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1
	_url := fmt.Sprintf("%v%v", base, bfes.binanceFutureUrl.SetLeverage)
	request.PrepareRequest()

	bnreq := bncommon.NewBinanceServiceHttpRequest[bnserivcemodelreq.SetLeverageBinanceServiceRequest]()
	err := bnreq.NewBinanceHttpRequest(_url)
	if err != nil {
		return err
	}

	signature := bnreq.CreateRequestSignUrl(request, bfes.secrets.BinanceSecretKey)
	bnreq.RequestPostMethod(signature)
	bnreq.AddHeader(bfes.secrets.BinanceApiKey)

	bntransport := bncommon.NewBinanceTransport(&http.Transport{})
	bnclient := bncommon.NewBinanceSerivceHttpClient(bntransport.GetTransport())
	err = bnclient.Do(bnreq.GetBinanceRequest())
	if err != nil {
		return err
	}

	bnres := bnclient.GetBinanceHttpClientResponse()
	if bnres.StatusCode != http.StatusOK {
		bnResponseError := new(bnservicemodelres.ResponseBinanceFutureError)
		json.NewDecoder(bnres.Body).Decode(bnResponseError)
		msg := common.FormatMessageOtherThanHttpStatus200(
			bfes.binanceFutureServiceName,
			bnres.StatusCode,
			bnResponseError.Code,
			bnResponseError.Message,
		)
		return errors.New(msg)
	}

	return nil
}
