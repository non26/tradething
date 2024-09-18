package bnfuture

type ResponseBinanceFutureError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
