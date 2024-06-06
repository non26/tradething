package handler

import (
	"net/http"
	model "tradething/app/bn/app/model/handlermodel/future"
	service "tradething/app/bn/app/service/future"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IPlaceSingleOrderHandler interface {
	GetRequestBody(c echo.Context) (*model.PlaceSignleOrderHandlerRequest, error)
	Handler(c echo.Context) error
}

type placeSinglerOrderHandler struct {
	service service.IBinanceFutureService
}

func NewPlaceSinglerOrderHandler(
	service service.IBinanceFutureService,
) IPlaceSingleOrderHandler {
	return &placeSinglerOrderHandler{
		service,
	}
}

func (h *placeSinglerOrderHandler) GetRequestBody(
	c echo.Context,
) (*model.PlaceSignleOrderHandlerRequest, error) {
	req := new(model.PlaceSignleOrderHandlerRequest)
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
