package handler

import (
	"net/http"
	model "tradething/app/bn/app/model/handlermodel/future"
	service "tradething/app/bn/app/service/future"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type ISetNewLeveragehandler interface {
	GetRequestBody(c echo.Context) (*model.SetLeverageHandlerRequest, error)
	Handler(c echo.Context) error
}

type setNewLeveragehandler struct {
	service service.IBinanceFutureService
}

func NewsetNewLeveragehandler(
	service service.IBinanceFutureService,
) ISetNewLeveragehandler {
	return &setNewLeveragehandler{
		service,
	}
}

func (h *setNewLeveragehandler) GetRequestBody(c echo.Context) (*model.SetLeverageHandlerRequest, error) {
	req := new(model.SetLeverageHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *setNewLeveragehandler) Handler(c echo.Context) error {

	request, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
				Data:    nil,
			},
		)
	}

	err = h.service.SetNewLeverage(
		c.Request().Context(),
		request,
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
				Data:    nil,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		common.CommonResponse{
			Code:    common.SuccessCode,
			Message: "success",
			Data:    nil,
		},
	)
}
