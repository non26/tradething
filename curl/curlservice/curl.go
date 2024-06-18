package curlservice

type ICurl interface {
	SetUrl(url string)
	SetMethod(method string)
	PrepareCurl() ICurl
	ExecuteCurl() error
	SetBody(body interface{})
}
