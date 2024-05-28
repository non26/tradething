package model

type CommonOKXServiceResponse struct {
	Code    string        `json:"code"`
	Message string        `json:"msg"`
	Data    []interface{} `json:"data"`
}
