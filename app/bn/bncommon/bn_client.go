package bncommon

import "net/http"

type binanceSerivceHttpClient struct {
	client   *http.Client
	response *http.Response
}

func NewBinanceSerivceHttpClient(
	t *http.Transport,
) *binanceSerivceHttpClient {
	binanceClient := binanceSerivceHttpClient{}
	binanceClient.client = &http.Client{
		Transport: t,
	}
	return &binanceClient
}

func (b *binanceSerivceHttpClient) Do(bnreq *http.Request) error {
	res, err := b.client.Do(bnreq)
	if err != nil {
		return err
	}
	b.response = res
	return nil
}

func (b *binanceSerivceHttpClient) GetBinanceHttpClientResponse() *http.Response {
	return b.response
}
