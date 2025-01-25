package bnfuture

import (
	"net/http"
	handlerreq "tradething/app/bn/bn_future/handler_request"
	bnfuture "tradething/app/bn/bn_future/service"

	"github.com/labstack/echo/v4"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type ISetPositionHandler interface {
	GetRequestBody(c echo.Context) (*handlerreq.PlacePosition, error)
	Handler(c echo.Context) error
}

type setPositionHandler struct {
	service bnfuture.IBinanceFutureService
}

func NewSetPositionHandler(
	service bnfuture.IBinanceFutureService,
) ISetPositionHandler {
	return &setPositionHandler{
		service,
	}
}

func (h *setPositionHandler) GetRequestBody(c echo.Context) (*handlerreq.PlacePosition, error) {
	req := new(handlerreq.PlacePosition)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *setPositionHandler) Handler(c echo.Context) error {
	req, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: err.Error(),
		})
	}

	response, svcerr := h.service.SetAdvancedPosition(c.Request().Context(), req.ToServiceModel())
	if svcerr != nil {
		return c.JSON(http.StatusInternalServerError, utils.CommonResponse{
			Code:    utils.FailCode,
			Message: svcerr.Error(),
		})
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
