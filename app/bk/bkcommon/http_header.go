package bkcommon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"tradetoolv2/common"
)

func CreateBitkubHeaders(
	req *http.Request,
	apiKey string,
	timestamp int64,
	sign string,
) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-BTK-TIMESTAMP", fmt.Sprintf("%v", timestamp))
	req.Header.Add("X-BTK-APIKEY", apiKey)
	req.Header.Add("X-BTK-SIGN", sign)
}

func CreateSignUrlMethod(
	method string,
	timestamp int64,
	urlPath string,
	bodyJson []byte,
	sk string,
	params map[string]string,
) string {
	payloadString := ""
	if method == http.MethodPost {
		if bodyJson != nil {
			payloadString = common.JsonToString(bodyJson)
		}
	} else {
		if params != nil {
			payloadString = common.ParamToQueryString(params)
		}
	}
	s := fmt.Sprintf("%v%v%v%v", timestamp, method, urlPath, payloadString)
	mac := hmac.New(sha256.New, []byte(sk))
	_, _ = mac.Write([]byte(s))
	stringOfHex := hex.EncodeToString(mac.Sum(nil))
	return stringOfHex
}
