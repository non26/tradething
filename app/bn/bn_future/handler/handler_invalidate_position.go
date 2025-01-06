package bnfuture

import (
	"net/http"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
)

type IInvalidatePositionHandler interface {
	GetRequestBody(c echo.Context) (*bnfuturereq.InvalidatePositionHandlerRequest, error)
	Handler(c echo.Context) error
}

type invalidatePositionHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewInvalidatePositionHandler(
	service bnfuture.IBinanceFutureService,
) IInvalidatePositionHandler {
	return &invalidatePositionHandler{
		service,
	}
}

func (h *invalidatePositionHandler) GetRequestBody(c echo.Context) (*bnfuturereq.InvalidatePositionHandlerRequest, error) {
	req := new(bnfuturereq.InvalidatePositionHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *invalidatePositionHandler) Handler(c echo.Context) error {
	req, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.InvalidatePosition(c.Request().Context(), req.ToServiceRequest())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
