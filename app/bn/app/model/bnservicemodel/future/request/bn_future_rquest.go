package bnserivcemodelreq

type IBnFutureServiceRequest interface {
	PrepareRequest()
	GetData() interface{}
}
