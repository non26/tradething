package bnfuture

import (
	"net/http"
	bnfuturereq "tradething/app/bn/bn_future/handler_request_model"
	bnfuture "tradething/app/bn/bn_future/service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IqueryOrderHandler interface {
	GetRequestBody(c echo.Context) (*bnfuturereq.QueryOrderBinanceHandlerRequest, error)
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
) (*bnfuturereq.QueryOrderBinanceHandlerRequest, error) {
	req := new(bnfuturereq.QueryOrderBinanceHandlerRequest)
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
			common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
				Data:    nil,
			},
		)
	}
	request.Transform()

	res, err := h.service.QueryOrder(
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
