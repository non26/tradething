package bkservice

import (
	"context"
	"net/http"
	model "tradething/app/bk/app/model/bitkubservicemodel"
	"tradething/app/bk/bkcommon"
)

func (o *orderBkService) BuyOrder(
	ctx context.Context,
	req *model.BuyOrderBkServiceRequest) (*model.BuyOrderBkServiceResponse, error) {
	bkRequest := bkcommon.NewHttpBitkubRequest[*model.BuyOrderBkServiceRequest](
		o.serviceName,
		http.MethodPost,
		o.kubConfigUrl.BaseUrl,
		o.kubConfigUrl.BuyOrderUrl,
		req,
	)
	bkRequest, err := bkRequest.GetBkNewRequest().
		SetBkHeaders(o.bkApiKey, o.bkSecretKey).
		Error()
	if err != nil {
		return nil, err
	}

	bkResponse := bkcommon.NewHttpBitkubResponse[*model.BuyOrderBkServiceResponse, *model.BuyOrderBkServiceRequest](
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
