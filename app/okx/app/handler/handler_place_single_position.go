package handler

import (
	"net/http"
	handlermodel "tradetoolv2/app/okx/app/model/handlermodel"
	"tradetoolv2/app/okx/app/service"
	"tradetoolv2/common"

	"github.com/labstack/echo/v4"
)

type IPlaceSinglePositionHandler interface {
	Handler(c echo.Context) error
}

type placeSinglePositionHandler struct {
	service service.IOkxAppService
}

func NewPlaceSinglePositionHandler(
	service service.IOkxAppService,
) IPlaceSinglePositionHandler {
	return &placeSinglePositionHandler{
		service,
	}
}

func (p *placeSinglePositionHandler) Handler(c echo.Context) error {
	body := new(handlermodel.PlaceASinglePositionHandlerRequest)
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

	data, err := p.service.PlaceAPosition(body)
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
		&common.CommonResponse{
			Code:    common.SuccessCode,
			Message: "Success",
			Data:    data,
		},
	)
}
