package bnfuture

import (
	"net/http"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfuture "tradething/app/bn/bn_future/service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IPlaceSingleOrderHandler interface {
	GetRequestBody(c echo.Context) (*bnfuturereq.PlaceSignleOrderHandlerRequest, error)
	Handler(c echo.Context) error
}

type placeSinglerOrderHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewPlaceSinglerOrderHandler(
	service bnfuture.IBinanceFutureService,
) IPlaceSingleOrderHandler {
	return &placeSinglerOrderHandler{
		service,
	}
}

func (h *placeSinglerOrderHandler) GetRequestBody(
	c echo.Context,
) (*bnfuturereq.PlaceSignleOrderHandlerRequest, error) {
	req := new(bnfuturereq.PlaceSignleOrderHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *placeSinglerOrderHandler) Handler(c echo.Context) error {

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

	res, err := h.service.PlaceSingleOrder(
		c.Request().Context(),
		request.ToServiceModel(),
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
