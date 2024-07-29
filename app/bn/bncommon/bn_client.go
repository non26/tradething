package bncommon

import "net/http"

type IBinanceSerivceHttpClient interface {
	Do(bnreq *http.Request) error
	GetBinanceHttpClientResponse() *http.Response
	SetClient(t *http.Transport)
}

type binanceSerivceHttpClient struct {
	client   *http.Client
	response *http.Response
}

func NewBinanceSerivceHttpClient() IBinanceSerivceHttpClient {
	binanceClient := binanceSerivceHttpClient{}
	return &binanceClient
}

func (b *binanceSerivceHttpClient) SetClient(t *http.Transport) {
	b.client = &http.Client{
		Transport: t,
	}
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
