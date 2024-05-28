package okxservice

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	model "tradetoolv2/app/okx/app/model/okxservicemodel"
	"tradetoolv2/app/okx/okxcommon"
)

func (o *okxExternalService) PlaceMultiplePosition(
	body *model.PlaceMultiplePositionOKXServiceRequest,
) (*http.Response, error) {
	_endPoint := o.okxFutureUrl.PlaceMultiPosition
	_url := fmt.Sprintf("%v%v", o.okxFutureUrl.OkxFutureBaseUrl.Okx1, _endPoint)
	_method := http.MethodPost
	_body, err := okxcommon.StructToJson(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(_method, _url, bytes.NewReader(_body))
	if err != nil {
		return nil, errors.New("OKX-PlaceMultipleOrder Request Error: " + err.Error())
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
		return nil, errors.New("OKX-PlaceMultipleOrder Response Error: " + err.Error())
	}

	return res, nil
}
