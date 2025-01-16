package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request_model"
	bnfuture "tradething/app/bn/bn_future/service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IPlaceMultipleOrderHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.PlaceMultiplePositions, error)
	Handler(c echo.Context) error
}

type placeMultipleOrderHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewPlaceMultipleOrderHandler(
	service bnfuture.IBinanceFutureService,
) IPlaceMultipleOrderHandler {
	return &placeMultipleOrderHandler{
		service,
	}
}

func (h *placeMultipleOrderHandler) GetRequestBody(c echo.Context) (*handlerreq.PlaceMultiplePositions, error) {
	req := new(handlerreq.PlaceMultiplePositions)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *placeMultipleOrderHandler) Handler(c echo.Context) error {
	request, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.CommonResponse{
			Code:    common.FailCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	request.Transform()
	response, err := h.service.PlaceMultiOrder(c.Request().Context(), request.ToServiceModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.CommonResponse{
			Code:    common.FailCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.CommonResponse{
		Code:    common.SuccessCode,
		Message: "success",
		Data:    response,
	})
}
