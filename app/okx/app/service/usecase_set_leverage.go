package service

import (
	handlermodel "tradething/app/okx/app/model/handlermodel"
)

func (o *okxAppService) SetNewLeverage(req *handlermodel.SetNewLeverageHandlerRequest) (*handlermodel.SetNewLeverageHandlerResponse, error) {
	data, err := o.okxExtService.SetLeverage(
		req.ToSetNewLeverageOKXServiceRequest(),
	)
	if err != nil {
		return nil, err
	}

	res := &handlermodel.SetNewLeverageHandlerResponse{}
	res.ToSetNewLeverageHandlerResponse(data)

	return res, nil
}
