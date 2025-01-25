package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type ICloseByClientIdsHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.ClosePositionByClientIds, error)
	Handler(c echo.Context) error
}

type closeByClientIdsHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewCloseByClientIdsHandler(
	service bnfuture.IBinanceFutureService,
) ICloseByClientIdsHandler {
	return &closeByClientIdsHandler{
		service,
	}
}

func (h *closeByClientIdsHandler) GetRequestBody(c echo.Context) (*handlerreq.ClosePositionByClientIds, error) {
	req := new(handlerreq.ClosePositionByClientIds)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *closeByClientIdsHandler) Handler(c echo.Context) error {
	request, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}
	response, err := h.service.CloseByClientIds(c.Request().Context(), request.ToServiceModel())
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
