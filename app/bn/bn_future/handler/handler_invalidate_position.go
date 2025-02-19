package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type IInvalidatePositionHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.InvalidatePosition, error)
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

func (h *invalidatePositionHandler) GetRequestBody(c echo.Context) (*handlerreq.InvalidatePosition, error) {
	req := new(handlerreq.InvalidatePosition)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *invalidatePositionHandler) Handler(c echo.Context) error {
	req, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}

	response, err := h.service.InvalidatePosition(c.Request().Context(), req.ToServiceModel())
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
