package bkservice

import (
	"context"
	"net/http"
	model "tradetoolv2/app/bk/app/model/bitkubservicemodel"
	"tradetoolv2/app/bk/bkcommon"
)

func (o *orderBkService) SellOrder(
	ctx context.Context,
	req *model.SellBkServiceRequest) (*model.SellOrderBkServiceResponse, error) {
	bkRequest := bkcommon.NewHttpBitkubRequest[*model.SellBkServiceRequest](
		o.serviceName,
		http.MethodPost,
		o.kubConfigUrl.BaseUrl,
		o.kubConfigUrl.SellOrderUrl,
		req,
	)
	bkRequest, err := bkRequest.GetBkNewRequest().
		SetBkHeaders(o.bkApiKey, o.bkSecretKey).
		Error()
	if err != nil {
		return nil, err
	}

	bkResponse := bkcommon.NewHttpBitkubResponse[*model.SellOrderBkServiceResponse, *model.SellBkServiceRequest](
		bkRequest,
		o.serviceName,
	)

	bkResponse, err = bkResponse.SetTransport().
		SetClient().
		GetNewBkResponse().
		GetBkResponseBody().
		Error()
	if err != nil {
		return nil, err
	}
	return bkResponse.GetBody(), nil
}
