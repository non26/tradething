package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type IPlaceSingleOrderHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.PlacePosition, error)
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
) (*handlerreq.PlacePosition, error) {
	req := new(handlerreq.PlacePosition)
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
			utils.CommonResponse{
				Code:    utils.FailCode,
				Message: err.Error(),
			},
		)
	}
	request.Transform()
	err = request.Validate()
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			utils.CommonResponse{
				Code:    utils.FailCode,
				Message: err.Error(),
			},
		)
	}

	response, err := h.service.PlaceSingleOrder(
		c.Request().Context(),
		request.ToServiceModel(),
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			utils.CommonResponse{
				Code:    utils.FailCode,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		utils.CommonResponse{
			Code:    utils.SuccessCode,
			Message: "success",
			Data:    response,
		},
	)
}
