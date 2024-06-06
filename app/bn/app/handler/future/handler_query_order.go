package handler

import (
	"net/http"
	model "tradething/app/bn/app/model/handlermodel/future"
	service "tradething/app/bn/app/service/future"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IqueryOrderHandler interface {
	GetRequestBody(c echo.Context) (*model.QueryOrderBinanceHandlerRequest, error)
	Handler(c echo.Context) error
}

type queryOrderHandler struct {
	service service.IBinanceFutureService
}

func NewqueryOrderHandler(
	service service.IBinanceFutureService,
) IqueryOrderHandler {
	return &queryOrderHandler{
		service,
	}
}

func (h *queryOrderHandler) GetRequestBody(
	c echo.Context,
) (*model.QueryOrderBinanceHandlerRequest, error) {
	req := new(model.QueryOrderBinanceHandlerRequest)
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

	res, err := h.service.QueryOrder(
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
