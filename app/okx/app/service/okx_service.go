package service

import (
	handlermodel "tradetoolv2/app/okx/app/model/handlermodel"
	"tradetoolv2/app/okx/app/okxservice"
)

type IOkxAppService interface {
	PlaceAPosition(req *handlermodel.PlaceASinglePositionHandlerRequest) (*handlermodel.PlaceASinglePositionHandlerResponse, error)
	PlaceMultiplePosition(req *handlermodel.PlaceMultiplePositionHandlerRequest) error
	SetNewLeverage(req *handlermodel.SetNewLeverageHandlerRequest) (*handlermodel.SetNewLeverageHandlerResponse, error)
}

type okxAppService struct {
	okxExtService okxservice.IOKXExternalService
}

func NewOkxService(
	okxExtService okxservice.IOKXExternalService,
) IOkxAppService {
	return &okxAppService{
		okxExtService,
	}
}
