package handler

import (
	"net/http"
	model "tradething/app/kc/app/model/handlermodel/future/order"
	service "tradething/app/kc/app/service/future/order"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type placeFutureOrderHandler struct {
	orderService service.IFutureOrderService
}

func NewPlaceSpotORderHandler(
	orderService service.IFutureOrderService,
) *placeFutureOrderHandler {
	return &placeFutureOrderHandler{
		orderService: orderService,
	}
}

func (p *placeFutureOrderHandler) Handler(e echo.Context) error {
	// get body
	body := new(model.PlaceFutureOrderHandlerRequest)
	err := e.Bind(body)
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			common.CommonResponse{
				Code:    common.FailCode,
				Message: err.Error(),
			},
		)
	}

	_, err = p.orderService.PlaceFutureOrderService(
		e.Request().Context(),
		body,
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
