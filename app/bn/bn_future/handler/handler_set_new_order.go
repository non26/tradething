package bnfuture

import (
	"net/http"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfuture "tradething/app/bn/bn_future/service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type ISetNewLeveragehandler interface {
	GetRequestBody(c echo.Context) (*bnfuturereq.SetLeverageHandlerRequest, error)
	Handler(c echo.Context) error
}

type setNewLeveragehandler struct {
	service bnfuture.IBinanceFutureService
}

func NewsetNewLeveragehandler(
	service bnfuture.IBinanceFutureService,
) ISetNewLeveragehandler {
	return &setNewLeveragehandler{
		service,
	}
}

func (h *setNewLeveragehandler) GetRequestBody(c echo.Context) (*bnfuturereq.SetLeverageHandlerRequest, error) {
	req := new(bnfuturereq.SetLeverageHandlerRequest)
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

	res, err := h.service.SetNewLeverage(
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
			Data:    res,
		},
	)
}
