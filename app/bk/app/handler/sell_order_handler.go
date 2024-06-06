package handler

import (
	"net/http"
	model "tradething/app/bk/app/model/handlermodel"
	service "tradething/app/bk/app/service/spot/order"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type sellOrderHandler struct {
	orederService service.IOrderService
}

func NewSellOrderHandler(
	orederService service.IOrderService,
) *sellOrderHandler {
	return &sellOrderHandler{
		orederService: orederService,
	}
}

func (b *sellOrderHandler) Handler(e echo.Context) error {
	// get request body
	req := new(model.SellHandlerRequest)
	err := e.Bind(req)
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
			},
		)
	}
	_, err = b.orederService.SellOrder(
		e.Request().Context(),
		req,
	)
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
			},
		)
	}
	return e.JSON(
		http.StatusOK,
		common.CommonResponse{
			Code:    common.SuccessCode,
			Message: "success",
		},
	)
}
