package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type IAdvancedPositionHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.AdvancedPosition, error)
	Handler(c echo.Context) error
}

type advancedPositionHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewAdvancedPositionHandler(
	service bnfuture.IBinanceFutureService,
) IAdvancedPositionHandler {
	return &advancedPositionHandler{
		service,
	}
}

func (a *advancedPositionHandler) GetRequestBody(c echo.Context) (*handlerreq.AdvancedPosition, error) {
	req := new(handlerreq.AdvancedPosition)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (a *advancedPositionHandler) Handler(c echo.Context) error {
	req, err := a.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}

	response, err := a.service.SetAdvancedPosition(c.Request().Context(), req.ToPosition())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, utils.CommonResponse{
		Code:    utils.SuccessCode,
		Message: "success",
		Data:    response,
	})
}
