package handlers

import (
	"net/http"
	req "tradething/app/bn/handlers/spot/req"
	process "tradething/app/bn/process/spot"

	"github.com/labstack/echo/v4"
)

type IMultipleOrderHandler interface {
	Handler(c echo.Context) error
}

type multipleOrderHandler struct {
	processMultipleOrder process.ISpot
}

func NewMultipleOrderHandler(
	processMultipleOrder process.ISpot,
) IMultipleOrderHandler {
	return &multipleOrderHandler{processMultipleOrder}
}

func (h *multipleOrderHandler) Handler(c echo.Context) error {
	request := req.MultipleOrders{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.processMultipleOrder.MultiplePosition(c.Request().Context(), request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
