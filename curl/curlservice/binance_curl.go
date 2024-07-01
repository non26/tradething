package curlservice

import (
	"net/http"
	"tradething/app/bn/bncommon"
	"tradething/config"
	"tradething/curl/constant"
	"tradething/curl/curlcommon"
)

type binanceCurl struct {
	header        map[string]string
	method        string
	_url          string
	curl_command  []string
	post_req_body string
}

func NewBinanceCurl(config *config.AppConfig) ICurl {
	header := make(map[string]string)
	header[constant.Bn_Header_BMX] = config.Secrets.BinanceApiKey
	header[constant.Bn_Header_Content_Type] = constant.Bn_Content_Type
	b := &binanceCurl{}
	b.header = header
	b.curl_command = []string{}
	return b
}

func (b *binanceCurl) PrepareCurl() ICurl {

	b.curl_command = append(b.curl_command, "curl")
	b.curl_command = append(b.curl_command, "-H")
	b.curl_command = append(b.curl_command, curlcommon.MatchHeader(constant.Bn_Header_BMX, b.header[constant.Bn_Header_BMX]))
	b.curl_command = append(b.curl_command, "-H")

	b.curl_command = append(b.curl_command, curlcommon.MatchHeader(constant.Bn_Header_Content_Type, b.header[constant.Bn_Header_Content_Type]))
	switch b.method {
	case http.MethodPost:
		b.curl_command = append(b.curl_command, "-X")
		b.curl_command = append(b.curl_command, http.MethodPost)
		b.curl_command = append(b.curl_command, "-d")
		b.curl_command = append(b.curl_command, b.post_req_body)
		b.curl_command = append(b.curl_command, b._url)
		b.curl_command = append(b.curl_command, "-k")
	case http.MethodGet:
		b.curl_command = append(b.curl_command, b._url+"?"+b.post_req_body)
		b.curl_command = append(b.curl_command, "-k")
	}
	return b
}

func (b *binanceCurl) ExecuteCurl() error {
	err := curlcommon.ExcecuteCurl(b.curl_command)
	if err != nil {
		return err
	}
	return nil
}

func (b *binanceCurl) SetMethod(method string) {
	b.method = method
}

func (b *binanceCurl) SetUrl(url string) {
	b._url = url
}

func (b *binanceCurl) SetBody(body interface{}) {
	b.post_req_body = body.(string)
}

func SetBinanceSignUrl[T any](m T, binance_secret_key string) string {
	_url_value := bncommon.GetQueryStringFromStructType(m)
	sign := bncommon.CreateBinanceSignature(_url_value, binance_secret_key)
	return sign
}
