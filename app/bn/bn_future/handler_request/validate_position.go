package bnfuture

import model "tradething/app/bn/bn_future/service_model"

type ValidatePosition struct {
	ClientIds []string `json:"client_ids"`
}

func (v *ValidatePosition) ToServiceModel() *model.ClientIds {
	serviceModel := &model.ClientIds{}
	serviceModel.SetClientIds(v.ClientIds)
	return serviceModel
}
