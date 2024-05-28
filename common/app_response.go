package common

import (
	"fmt"
	"strings"
)

var SuccessCode string = "0000"
var FailCode string = "9999"

var _format_message = `<serviceName>-<httpStatus>-<serviceNameCode>-<serviceNameMessage>`

func FormatMessageOtherThanHttpStatus200(
	serviceName string,
	httpStatus int,
	serviceNameCode int,
	serviceNameMessage string,
) string {
	_f := strings.Replace(_format_message, "<serviceName>", serviceName, -1)
	_f = strings.Replace(_f, "<httpStatus>", fmt.Sprintf("%v", httpStatus), -1)
	_f = strings.Replace(_f, "<serviceNameCode>", fmt.Sprintf("%v", serviceNameCode), -1)
	_f = strings.Replace(_f, "<serviceNameMessage>", serviceNameMessage, -1)
	return _f
}

type CommonResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
