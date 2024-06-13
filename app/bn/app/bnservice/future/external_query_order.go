package bnservice

import (
	"context"
	"fmt"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/app/bn/bncommon"
)

// func (bfes *binanceFutureExternalService) QueryOrder(
// 	ctx context.Context,
// 	request *bnserivcemodelreq.QueryOrderBinanceServiceRequest,
// ) (*bnservicemodelres.QueryOrderBinanceServiceResponse, error) {
// 	endPoint := bfes.binanceFutureUrl.QueryOrder
// 	_url := fmt.Sprintf("%v%v", bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, endPoint)

// 	t := time.Now().Unix() * 1000
// 	tt := strconv.FormatInt(t, 10)
// 	data := url.Values{}
// 	data.Set("symbol", request.Symbol)
// 	data.Set("timestamp", tt)
// 	data.Set("origClientOrderId", request.Symbol)
// 	encodeData := bncommon.CreateBinanceSignature(data, bfes.secrets.BinanceSecretKey)

// 	req, err := http.NewRequest(http.MethodGet, _url, nil)
// 	// req, err := http.NewRequest(http.MethodGet, _url, strings.NewReader(encodeData))
// 	if err != nil {
// 		return nil, errors.New("Query Order Request Error: " + err.Error())
// 	}

// 	req.URL.RawQuery = encodeData

// 	req.Header.Add("X-MBX-APIKEY", bfes.secrets.BinanceApiKey)
// 	req.Header.Add("CONTENT-TYPE", "application/x-www-form-urlencoded")

// 	transport := &http.Transport{}
// 	client := &http.Client{
// 		Transport: transport,
// 	}
// 	var res *http.Response
// 	res, err = client.Do(req)
// 	if err != nil {
// 		return nil, errors.New("Query Order Response Error: " + err.Error())
// 	}

// 	var m *bnservicemodelres.QueryOrderBinanceServiceResponse
// 	json.NewDecoder(res.Body).Decode(m)

// 	return (*bnservicemodelres.QueryOrderBinanceServiceResponse)(m.ToHandlerResponse()), nil
// }

func (bfes *binanceFutureExternalService) QueryOrder(
	ctx context.Context,
	request *bnserivcemodelreq.QueryOrderBinanceServiceRequest,
) (*bnservicemodelres.QueryOrderBinanceServiceResponse, error) {

	endPoint := bfes.binanceFutureUrl.QueryOrder
	_url := fmt.Sprintf("%v%v", bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, endPoint)
	request.PrepareRequest()

	bnreq := bncommon.NewBinanceServiceHttpRequest[bnserivcemodelreq.QueryOrderBinanceServiceRequest]()
	err := bnreq.NewBinanceHttpRequest(_url)
	if err != nil {
		return nil, err
	}

	signature := bnreq.CreateRequestSignUrl(request, bfes.secrets.BinanceSecretKey)
	bnreq.RequestGetMethod(signature)
	bnreq.AddHeader(bfes.secrets.BinanceApiKey)

	bntransport := bncommon.NewBinanceTransport(&http.Transport{})
	bnclient := bncommon.NewBinanceSerivceHttpClient(bntransport.GetTransport())
	err = bnclient.Do(bnreq.GetBinanceRequest())
	if err != nil {
		return nil, err
	}

	bnres := bncommon.NewBinanceServiceHttpResponse[bnservicemodelres.QueryOrderBinanceServiceResponse](
		bnclient.GetBinanceHttpClientResponse(),
	)
	err = bnres.DecodeBinanceServiceResponse(bfes.binanceFutureServiceName)
	if err != nil {
		return nil, err
	}

	return bnres.GetBinanceServiceResponse(), nil
}
