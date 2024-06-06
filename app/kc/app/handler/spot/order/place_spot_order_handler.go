package handelr

import (
	"net/http"
	model "tradething/app/kc/app/model/handlermodel/spot/order"
	service "tradething/app/kc/app/service/spot/order"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type placeSpotOrderHandler struct {
	orderService service.ISpotOrderService
}

func NewPlaceSpotORderHandler(
	orderService service.ISpotOrderService,
) *placeSpotOrderHandler {
	return &placeSpotOrderHandler{
		orderService: orderService,
	}
}

func (p *placeSpotOrderHandler) Handler(e echo.Context) error {
	// get body
	body := new(model.PlaceSpotOrderHandlerRequest)
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

	_, err = p.orderService.PlaceSpotOrderService(
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
