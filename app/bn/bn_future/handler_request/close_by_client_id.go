package bnfuture

import model "tradething/app/bn/bn_future/service_model"

type ClosePositionByClientIds struct {
	ClientIds []string `json:"client_ids"`
}

func (c *ClosePositionByClientIds) ToServiceModel() *model.ClientIds {
	serviceModel := &model.ClientIds{}
	serviceModel.SetClientIds(c.ClientIds)
	return serviceModel
}
