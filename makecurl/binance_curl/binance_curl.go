package binancecurl

import (
	"net/http"
	"tradething/curl/curlcommon"
	"tradething/makecurl/constants"
)

// SetMethod
// PrepareCurl

type BinaceFutureCurl struct {
	base_url     string
	method       string
	curl_command []string
}

func NewBinanceFutureCurl() *BinaceFutureCurl {
	_b := BinaceFutureCurl{
		base_url: "",
	}

	return &_b
}

func (b *BinaceFutureCurl) SetMethod(method string) {
	b.method = method
}

func (b *BinaceFutureCurl) PrepareCurl() {

	b.curl_command = append(b.curl_command, "curl")
	b.curl_command = append(b.curl_command, "-H")
	// b.curl_command = append(b.curl_command, curlcommon.MatchHeader(constant.Bn_Header_BMX, b.header[constant.Bn_Header_BMX]))
	b.curl_command = append(b.curl_command, curlcommon.MatchHeader(constants.Bn_Header_BMX, "Binance_api_Key"))
	b.curl_command = append(b.curl_command, "-H")

	// b.curl_command = append(b.curl_command, curlcommon.MatchHeader(constant.Bn_Header_Content_Type, b.header[constant.Bn_Header_Content_Type]))
	b.curl_command = append(b.curl_command, curlcommon.MatchHeader(constants.Bn_Content_Type, constants.Bn_Content_Type))
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

}
