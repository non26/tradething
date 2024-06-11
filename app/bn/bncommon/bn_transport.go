package bncommon

import "net/http"

type binanceServiceHttpTransport struct {
	t *http.Transport
}

func NewBinanceTransport(
	t *http.Transport,
) *binanceServiceHttpTransport {
	b := binanceServiceHttpTransport{
		t,
	}
	return &b
}

func (b *binanceServiceHttpTransport) GetTransport() *http.Transport {
	return b.t
}
