package bnfuture

type IBnFutureServiceRequest interface {
	PrepareRequest()
	GetData() interface{}
}
