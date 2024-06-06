package okxcommon

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"tradething/config"
)

func GenerateHeaders(
	req *http.Request,
	method string,
	requestPath string,
	body string,
	env string,
	secrets *config.Secrets,
) *http.Request {
	t := GenerateTimeHeader()
	req.Header.Add("OK-ACCESS-KEY", GenerateKeyHeader(secrets.OkxApiKey))
	req.Header.Add("OK-ACCESS-SIGN", GenerateSignHeader(t, method, requestPath, body, secrets.OkxSecretKey))
	req.Header.Add("OK-ACCESS-TIMESTAMP", t)
	req.Header.Add("OK-ACCESS-PASSPHRASE", GeneratePassPhaseHeader(secrets.OkxPassPhase))
	req.Header.Add("expTime", fmt.Sprint(time.Now().Add(3*time.Second).UnixMilli()))
	switch method {
	case http.MethodPost:
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("accept", "application/json")
	default:
		break
	}
	if strings.ToUpper(env) == "LOCAL" {
		req.Header.Add("x-simulated-trading", "1")
	}
	return req
}
