package handlers

import (
	"net/http"
	"tradething/app/bn/handlers/spot/req"
	process "tradething/app/bn/process/spot"

	"github.com/labstack/echo/v4"
)

type IOrderHandler interface {
	Handler(c echo.Context) error
}

type orderHandler struct {
	spotProcess process.ISpot
}

func NewOrderHandler(
	spotProcess process.ISpot,
) IOrderHandler {
	return &orderHandler{
		spotProcess: spotProcess,
	}
}

func (h *orderHandler) Handler(c echo.Context) error {
	request := req.Trade{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.spotProcess.PlaceOrder(c.Request().Context(), request.ToOrder())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
