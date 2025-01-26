package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	handlerres "tradething/app/bn/bn_future/handler_response"
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
		return c.JSON(http.StatusBadRequest, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}
	request.Transform()
	err = request.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}
	ctx := c.Request().Context()
	response := new(handlerres.PlacePosition)

	if len(request.InvalidatePosition) != 0 {
		invalidteResponse, err := h.service.InvalidatePosition(ctx, request.ToInvaldiatePositionServiceModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.CommonResponse{
				Code:    utils.FailCode,
				Message: err.Error(),
			})
		}
		response.InValidate = invalidteResponse
	}

	if len(request.ValidatePosition) != 0 {
		validateResponse, err := h.service.ValidateAdavancedPosition(ctx, request.ToValidatePosiionServiceModel())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.CommonResponse{
				Code:    utils.FailCode,
				Message: err.Error(),
			})
		}
		response.Validate = validateResponse
	}

	placeOrderResponse, err := h.service.PlaceSingleOrder(ctx, request.ToServiceModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}
	response.Symbol = placeOrderResponse.Symbol
	response.Quantity = placeOrderResponse.Quantity

	return c.JSON(
		http.StatusOK,
		utils.CommonResponse{
			Code:    utils.SuccessCode,
			Message: "success",
			Data:    response,
		})
}
