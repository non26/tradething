package bncommon

import "net/http"

func NewBinanceTransport(
	t *http.Transport,
) *http.Transport {
	return t
}
