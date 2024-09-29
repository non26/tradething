package apiinfo

import (
	requestbodyinfo "tradething/makecurl/registor_api/request_body_info"
)

type IApiInfo interface {
	SetRequestBody(field string, field_type string)
}

type apiInfo struct {
	api_alias    string
	api_endpoint string
	method       string
	requestBody  []requestbodyinfo.IRequestBodyInfo
}

func NewApi(
	api_alias string,
	api_endpoint string,
	method string,
) IApiInfo {
	return &apiInfo{
		api_alias,
		api_endpoint,
		method,
		make([]requestbodyinfo.IRequestBodyInfo, 0, 10),
	}
}

func (a *apiInfo) SetRequestBody(field string, field_type string) {
	_b := requestbodyinfo.NewRequestBodyInfo(field, field_type)
	a.requestBody = append(a.requestBody, _b)
}
