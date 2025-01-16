package bnfuture

type IBnFutureRequest interface {
	PrepareRequest()
	GetData() interface{}
}
