package handler

import (
	"net/http"
	model "tradetoolv2/app/bk/app/model/handlermodel"
	service "tradetoolv2/app/bk/app/service/spot/order"
	"tradetoolv2/common"

	"github.com/labstack/echo/v4"
)

type buyOrderHandler struct {
	orederService service.IOrderService
}

func NewBuyOrderHandler(
	orederService service.IOrderService,
) *buyOrderHandler {
	return &buyOrderHandler{
		orederService: orederService,
	}
}

func (b *buyOrderHandler) Handler(e echo.Context) error {
	// get request body
	req := new(model.BuyOrderHandlerRequest)
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
	_, err = b.orederService.BuyOrder(
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
