package bnservice

import (
	"context"
	"fmt"
	"net/http"
	bnserivcemodelreq "tradething/app/bn/app/model/bnservicemodel/future/request"
	bnservicemodelres "tradething/app/bn/app/model/bnservicemodel/future/response"
	"tradething/app/bn/bncommon"
)

// func (bfes *binanceFutureExternalService) 3PlaceSingleOrder(
// 	ctx context.Context,
// 	request *model.PlaceSignleOrderHandlerRequest,
// ) (*http.Response, error) {
// 	endPoint := bfes.binanceFutureUrl.SingleOrder
// 	_url := fmt.Sprintf("%v%v", bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, endPoint)
// 	request.ClientOrderId = request.Symbol
// 	t := time.Now().Unix() * 1000
// 	data := url.Values{}
// 	data.Set("symbol", request.Symbol)
// 	data.Set("side", request.Side)
// 	data.Set("positionSide", request.PositionSide)
// 	data.Set("type", "MARKET")
// 	data.Set("quantity", fmt.Sprintf("%v", request.EntryQuantity))
// 	data.Set("timestamp", strconv.FormatInt(t, 10))
// 	data.Set("newClientOrderId", request.ClientOrderId)
// 	encodeData := bncommon.CreateBinanceSignature(&data, bfes.secrets.BinanceSecretKey)

// 	req, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(encodeData))
// 	if err != nil {
// 		return nil, errors.New("PlaceSingleOrder Request Error: " + err.Error())
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
// 		return nil, errors.New("PlaceSingleOrder Response Error: " + err.Error())
// 	}

// 	return res, nil
// }

func (bfes *binanceFutureExternalService) PlaceSingleOrder(
	ctx context.Context,
	request *bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest,
) (*bnservicemodelres.PlaceSignleOrderBinanceServiceResponse, error) {
	endPoint := bfes.binanceFutureUrl.SingleOrder
	_url := fmt.Sprintf("%v%v", bfes.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1, endPoint)
	request.PrepareRequest()

	bnreq := bncommon.NewBinanceServiceHttpRequest[bnserivcemodelreq.PlaceSignleOrderBinanceServiceRequest]()
	err := bnreq.NewBinanceHttpRequest(_url)
	if err != nil {
		return nil, err
	}

	signature := bnreq.CreateRequestSignUrl(request, bfes.secrets.BinanceSecretKey)
	bnreq.RequestPostMethod(signature)
	bnreq.AddHeader(bfes.secrets.BinanceApiKey)

	bntransport := bncommon.NewBinanceTransport(&http.Transport{})
	bnclient := bncommon.NewBinanceSerivceHttpClient(bntransport.GetTransport())
	err = bnclient.Do(bnreq.GetBinanceRequest())
	if err != nil {
		return nil, err
	}

	bnres := bncommon.NewBinanceServiceHttpResponse[bnservicemodelres.PlaceSignleOrderBinanceServiceResponse](
		bnclient.GetBinanceHttpClientResponse())
	err = bnres.DecodeBinanceServiceResponse(bfes.binanceFutureServiceName)
	if err != nil {
		return nil, err
	}

	return bnres.GetBinanceServiceResponse(), nil
}
