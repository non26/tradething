package bnfuture

import svcfuture "tradething/app/bn/bn_future/service_model"

type InvalidatePositionHandlerRequest struct {
	OrderIds []string `json:"order_ids"`
}

func (r *InvalidatePositionHandlerRequest) ToServiceRequest() *svcfuture.InvalidatePositionServiceRequest {
	return &svcfuture.InvalidatePositionServiceRequest{
		OrderIds: r.OrderIds,
	}
}
