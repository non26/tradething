package kccommon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
)

type KucoinHeaderInfo struct {
	Request       *http.Request
	Method        string
	UrlPath       string
	Passphase     string
	SecretKey     string
	ApiKey        string
	ApiKeyVersion string
	KucoinTime    int64
	signature     string
}

func (k *KucoinHeaderInfo) createSignUrl(body string) {
	signature := fmt.Sprintf("%v%v%v%v", k.KucoinTime, k.Method, k.UrlPath, body)
	mac := hmac.New(sha256.New, []byte(k.SecretKey))
	_, _ = mac.Write([]byte(signature))
	signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	k.signature = signature
}

func (k *KucoinHeaderInfo) AddHeaders(body string) *http.Request {
	k.createSignUrl(body)
	k.Request.Header.Add("KC-API-KEY", k.ApiKey)
	k.Request.Header.Add("KC-API-SIGN", k.signature)
	k.Request.Header.Add("KC-API-TIMESTAMP", fmt.Sprintf("%v", k.KucoinTime))
	k.Request.Header.Add("KC-API-PASSPHRASE", k.Passphase)
	k.Request.Header.Add("KC-API-KEY-VERSION", k.ApiKeyVersion)
	return k.Request
}
