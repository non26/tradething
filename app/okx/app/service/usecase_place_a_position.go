package service

import (
	handlermodel "tradetoolv2/app/okx/app/model/handlermodel"
)

func (o *okxAppService) PlaceAPosition(req *handlermodel.PlaceASinglePositionHandlerRequest) (
	*handlermodel.PlaceASinglePositionHandlerResponse,
	error,
) {
	data, err := o.okxExtService.PlaceASinglePosition(req.ToPlaceASinglePositionOKXServiceRequest())
	if err != nil {
		return nil, err
	}
	res := &handlermodel.PlaceASinglePositionHandlerResponse{}
	res.ToPlaceASinglePositionHandlerRequest(data)
	return res, nil
}
