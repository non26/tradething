package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type IqueryOrderHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.QueryOrder, error)
	Handler(c echo.Context) error
}

type queryOrderHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewqueryOrderHandler(
	service bnfuture.IBinanceFutureService,
) IqueryOrderHandler {
	return &queryOrderHandler{
		service,
	}
}

func (h *queryOrderHandler) GetRequestBody(
	c echo.Context,
) (*handlerreq.QueryOrder, error) {
	req := new(handlerreq.QueryOrder)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *queryOrderHandler) Handler(c echo.Context) error {

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

	response, err := h.service.QueryOrder(
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
