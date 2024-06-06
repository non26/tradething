package handler

import (
	"net/http"
	handlermodel "tradething/app/okx/app/model/handlermodel"
	"tradething/app/okx/app/service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IPlaceMultiplePositionHandler interface {
	Handler(c echo.Context) error
}

type placeMultiplePositionHandler struct {
	service service.IOkxAppService
}

func NewPlaceMultiplePositionHandler(
	service service.IOkxAppService,
) IPlaceMultiplePositionHandler {
	return &placeMultiplePositionHandler{
		service,
	}
}

func (p *placeMultiplePositionHandler) Handler(c echo.Context) error {

	body := new(handlermodel.PlaceMultiplePositionHandlerRequest)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			&common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
			},
		)
	}

	err = p.service.PlaceMultiplePosition(body)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
			},
		)
	}
	return c.JSON(
		http.StatusOK,
		common.CommonResponse{
			Code:    common.SuccessCode,
			Message: "Success",
			Data:    nil,
		},
	)
}
