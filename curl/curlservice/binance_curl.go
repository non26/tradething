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
	return &binanceCurl{}
}

func (b *binanceCurl) PrepareCurl() ICurl {

	b.curl_command[0] = "curl"
	b.curl_command[1] = "-H"
	b.curl_command[2] = curlcommon.MatchHeader(constant.Bn_Header_BMX, b.header[constant.Bn_Header_BMX])
	b.curl_command[3] = "-H"
	b.curl_command[4] = curlcommon.MatchHeader(constant.Bn_Header_Content_Type, b.header[constant.Bn_Header_Content_Type])
	switch b.method {
	case http.MethodPost:
		b.curl_command[5] = "-X"
		b.curl_command[6] = http.MethodPost
		b.curl_command[7] = "-d"
		b.curl_command[8] = b.post_req_body
		b.curl_command[9] = b._url
		b.curl_command[10] = "-k"
	case http.MethodGet:
		b.curl_command[5] = b._url
		b.curl_command[6] = "-k"
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
	switch b.method {
	case b.method:
		b.post_req_body = body.(string)
	default:
		b.post_req_body = ""
	}
}

func SetBinanceSignUrl[T any](m T, binance_secret_key string) string {
	_url_value := bncommon.GetQueryStringFromStructType(m)
	sign := bncommon.CreateBinanceSignature(_url_value, binance_secret_key)
	return sign
}
