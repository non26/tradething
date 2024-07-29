package bncommon

import "net/http"

type IBinanceServiceHttpTransport interface {
	GetTransport() *http.Transport
}

type binanceServiceHttpTransport struct {
	t *http.Transport
}

func NewBinanceTransport(t *http.Transport) *binanceServiceHttpTransport {
	b := binanceServiceHttpTransport{
		t: t,
	}
	return &b
}

func (b *binanceServiceHttpTransport) GetTransport() *http.Transport {
	return b.t
}
