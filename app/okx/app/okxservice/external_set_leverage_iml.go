package okxservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	model "tradething/app/okx/app/model/okxservicemodel"
	"tradething/app/okx/okxcommon"
)

func (o *okxExternalService) SetLeverage(
	e *model.SetNewLeverageOKXServiceRequest,
) (*model.SetNewLeverageOKXserviceResponse, error) {

	body := model.SetNewLeverageOKXServiceRequest{}
	body.InstId = okxcommon.AddInstIdUSDTSWAPPostfix(body.InstId)

	_endPoint := o.okxFutureUrl.SetLeverage
	_url := fmt.Sprintf("%v%v", o.okxFutureUrl.OkxFutureBaseUrl.Okx1, _endPoint)
	_method := http.MethodPost
	_body, err := okxcommon.StructToJson(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(_method, _url, bytes.NewReader(_body))
	if err != nil {
		return nil, errors.New("OKX-SetLeverage Request Error: " + err.Error())
	}
	req = okxcommon.GenerateHeaders(
		req,
		_method,
		_endPoint,
		string(_body),
		o.env,
		o.secret,
	)

	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return nil, errors.New("OKX-SetLeverage Response Error: " + err.Error())
	}
	defer res.Body.Close()

	decodeResBody := &model.CommonOKXServiceResponse{}
	err = json.NewDecoder(res.Body).Decode(decodeResBody)
	if err != nil {
		return nil, errors.New("OKX-SetLeverage Decode Error: " + err.Error())
	}
	err = okxcommon.OkxConditionResponseError(res.StatusCode, decodeResBody.Code, decodeResBody.Message)
	if err != nil {
		return nil, err
	}

	eSetLeverage := []model.SetNewLeverageOKXserviceResponse{}
	eSetLeverage, err = okxcommon.ResponseToStruct[model.SetNewLeverageOKXserviceResponse](eSetLeverage, decodeResBody.Data)
	if err != nil {
		return nil, err
	}
	return &eSetLeverage[0], nil
}
