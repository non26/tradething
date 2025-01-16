package bnfuture

import model "tradething/app/bn/bn_future/service_model"

type InvalidatePosition struct {
	ClientIds []string `json:"client_ids"`
}

func (c *InvalidatePosition) ToServiceModel() *model.ClientIds {
	serviceModel := &model.ClientIds{}
	serviceModel.SetClientIds(c.ClientIds)
	return serviceModel
}
