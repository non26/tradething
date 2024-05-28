package bkservice

import (
	"context"
	"net/http"
	model "tradetoolv2/app/bk/app/model/bitkubservicemodel/account"
	"tradetoolv2/app/bk/bkcommon"
)

func (a *accountBkService) GetBalances(
	ctx context.Context,
) (*model.BalanceBkServiceResponse, error) {

	type EmptyPayload struct{}
	emptyPayload := EmptyPayload{}

	bkRequest := bkcommon.NewHttpBitkubRequest(
		a.serviceName,
		http.MethodPost,
		a.kubConfigUrl.BaseUrl,
		a.kubConfigUrl.Balances,
		emptyPayload,
	)

	bkRequest, err := bkRequest.GetBkNewRequest().
		SetBkHeaders(a.bkApiKey, a.bkSecretKey).
		Error()
	if err != nil {
		return nil, err
	}

	bkResponse := bkcommon.NewHttpBitkubResponse[*model.BalanceBkServiceResponse, EmptyPayload](
		bkRequest,
		a.serviceName,
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
