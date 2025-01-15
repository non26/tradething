package bnfuture

import svcfuture "tradething/app/bn/bn_future/service_model"

type CloseByClientIdHandlerRequest struct {
	ClientIds []string `json:"client_ids"`
}

func (c *CloseByClientIdHandlerRequest) ToServiceModel() *svcfuture.CloseByClientIdServiceRequest {
	serviceModel := &svcfuture.CloseByClientIdServiceRequest{}
	serviceModel.SetOrderIds(c.ClientIds)
	return serviceModel
}
